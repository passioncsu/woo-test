package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "admin123"

	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	assert.True(t, CheckPassword(password, hash))
	assert.False(t, CheckPassword("wrong-password", hash))
}
