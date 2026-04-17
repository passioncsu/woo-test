package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"student-admin/backend/internal/config"
	"student-admin/backend/internal/handler"
	"student-admin/backend/internal/middleware"
	"student-admin/backend/internal/repository"
	"student-admin/backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"student-admin/backend/internal/model"
)

func setupTestRouter(t *testing.T) (*gin.Engine, *config.Config) {
	gin.SetMode(gin.TestMode)

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Admin{}, &model.Student{})
	assert.NoError(t, err)

	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret:     "test-secret",
			Expiration: 24,
		},
	}

	adminRepo := repository.NewAdminRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	authService := service.NewAuthService(adminRepo, cfg)
	studentService := service.NewStudentService(studentRepo)
	authHandler := handler.NewAuthHandler(authService)
	studentHandler := handler.NewStudentHandler(studentService)

	r := gin.New()

	r.POST("/api/login", authHandler.Login)
	r.POST("/api/register", authHandler.Register)

	api := r.Group("/api")
	api.Use(middleware.JWTAuth(cfg))
	{
		api.GET("/students", studentHandler.List)
		api.POST("/students", studentHandler.Create)
		api.PUT("/students/:id", studentHandler.Update)
		api.DELETE("/students/:id", studentHandler.Delete)
	}

	return r, cfg
}

func TestRegisterAndLogin(t *testing.T) {
	r, _ := setupTestRouter(t)

	// 注册
	body, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": "admin123",
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// 登录
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	data := resp["data"].(map[string]interface{})
	token := data["accessToken"].(string)
	assert.NotEmpty(t, token)
}

func TestStudentCRUD(t *testing.T) {
	r, _ := setupTestRouter(t)

	// 先注册获取 token
	body, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": "admin123",
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	var loginResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &loginResp)
	token := loginResp["data"].(map[string]interface{})["accessToken"].(string)

	// 创建学生
	student := map[string]interface{}{
		"name":       "张三",
		"student_no": "2024001",
		"gender":     "男",
		"major":      "计算机科学",
	}
	body, _ = json.Marshal(student)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/students", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// 查询学生列表
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/students?keyword=张&page=1&pageSize=10", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var listResp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &listResp)
	data := listResp["data"].(map[string]interface{})
	assert.Equal(t, float64(1), data["total"])
}
