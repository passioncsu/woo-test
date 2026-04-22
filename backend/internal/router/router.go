package router

import (
	"student-admin/backend/internal/config"
	"student-admin/backend/internal/handler"
	"student-admin/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(
	engine *gin.Engine,
	cfg *config.Config,
	authHandler *handler.AuthHandler,
	studentHandler *handler.StudentHandler,
) {
	// 全局中间件
	engine.Use(middleware.Logger())
	engine.Use(middleware.CORS())
	engine.Use(gin.Recovery())

	// 健康检查
	engine.GET("/api/health", func(c *gin.Context) { c.Status(204) })

	// 公开接口
	engine.POST("/api/login", authHandler.Login)
	engine.POST("/api/register", authHandler.Register)
	engine.POST("/api/auth/logout", authHandler.Logout)

	// 需要鉴权的接口
	api := engine.Group("/api")
	api.Use(middleware.JWTAuth(cfg))
	{
		// 用户信息
		api.GET("/profile", authHandler.GetProfile)
		api.GET("/user/info", authHandler.GetProfile)

		// 权限码
		api.GET("/auth/codes", authHandler.GetAccessCodes)

		// 学生管理 CRUD
		api.GET("/students", studentHandler.List)
		api.GET("/students/:id", studentHandler.GetByID)
		api.POST("/students", studentHandler.Create)
		api.PUT("/students/:id", studentHandler.Update)
		api.DELETE("/students/:id", studentHandler.Delete)
	}
}
