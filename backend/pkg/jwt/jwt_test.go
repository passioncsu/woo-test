package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndParseToken(t *testing.T) {
	secret := "test-secret"
	userID := uint(1)
	username := "admin"
	expiration := 24

	// 生成 token
	token, err := GenerateToken(secret, userID, username, expiration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// 解析 token
	claims, err := ParseToken(secret, token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, username, claims.Username)
}

func TestParseToken_InvalidSecret(t *testing.T) {
	token, _ := GenerateToken("secret-a", 1, "admin", 24)

	// 用不同 secret 解析应失败
	_, err := ParseToken("secret-b", token)
	assert.Error(t, err)
}

func TestParseToken_InvalidToken(t *testing.T) {
	_, err := ParseToken("secret", "invalid-token-string")
	assert.Error(t, err)
}
