package middleware

import (
	"strings"

	"student-admin/backend/internal/config"
	"student-admin/backend/pkg/jwt"

	pkgresponse "student-admin/backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func JWTAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			pkgresponse.Unauthorized(c, "missing token")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			pkgresponse.Unauthorized(c, "invalid token format")
			return
		}

		claims, err := jwt.ParseToken(cfg.JWT.Secret, parts[1])
		if err != nil {
			pkgresponse.Unauthorized(c, "invalid or expired token")
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
