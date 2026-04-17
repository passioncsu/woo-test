package main

import (
	"fmt"
	"log"

	"student-admin/backend/internal/config"
	"student-admin/backend/internal/handler"
	"student-admin/backend/internal/model"
	"student-admin/backend/internal/repository"
	"student-admin/backend/internal/router"
	"student-admin/backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 连接数据库
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// 自动迁移
	if err := db.AutoMigrate(&model.Admin{}, &model.Student{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	// 初始化分层
	adminRepo := repository.NewAdminRepository(db)
	studentRepo := repository.NewStudentRepository(db)

	authService := service.NewAuthService(adminRepo, cfg)
	studentService := service.NewStudentService(studentRepo)

	authHandler := handler.NewAuthHandler(authService)
	studentHandler := handler.NewStudentHandler(studentService)

	// 初始化 Gin
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()

	// 注册路由
	router.Setup(engine, cfg, authHandler, studentHandler)

	// 启动服务
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := engine.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
