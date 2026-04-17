package handler

import (
	"student-admin/backend/internal/service"

	pkgresponse "student-admin/backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkgresponse.Fail(c, 400, 400, "invalid request: "+err.Error())
		return
	}

	resp, err := h.svc.Login(req)
	if err != nil {
		pkgresponse.Fail(c, 401, 401, err.Error())
		return
	}

	pkgresponse.OK(c, resp)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkgresponse.Fail(c, 400, 400, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Register(req); err != nil {
		pkgresponse.Fail(c, 400, 400, err.Error())
		return
	}

	pkgresponse.OKWithMessage(c, "register success", nil)
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	username, _ := c.Get("username")
	userID, _ := c.Get("user_id")

	pkgresponse.OK(c, gin.H{
		"user_id":  userID,
		"username": username,
	})
}

func (h *AuthHandler) GetAccessCodes(c *gin.Context) {
	pkgresponse.OK(c, []string{"*"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	pkgresponse.OKWithMessage(c, "logout success", nil)
}
