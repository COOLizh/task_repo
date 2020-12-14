package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/COOLizh/task_repo/pkg/models"
)

func TestCreateToken(t *testing.T) {
	user := models.User{
		ID:       102,
		Username: "John",
		Password: "Doe",
	}

	secretKey := "JWT_SALT"
	oldJwtSalt := os.Getenv(secretKey)
	defer os.Setenv(secretKey, oldJwtSalt)
	os.Setenv(secretKey, "verysecretsaltfortesting")

	actual, err := CreateToken(&user)
	assert.NoError(t, err, "test failed")
	t.Run("CreateToken", func(t *testing.T) {
		assert.NotEmpty(t, actual, "token string should not be empty")
		assert.NoError(t, err, "should be no errors")
	})
}

func TestParseToken(t *testing.T) {
	user := models.User{
		ID:       10,
		Username: "John",
		Password: "Doedoe",
	}

	secretKey := "JWT_SALT"
	oldJwtSalt := os.Getenv(secretKey)
	defer os.Setenv(secretKey, oldJwtSalt)
	os.Setenv(secretKey, "verysecretsaltfortesting")

	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAsImV4cCI6MTYzNDgwNjU4MX0.JJEmvlrdYPTv2lAeE3wSrPAjb188fl9QbC0l3hU7LLE"
	actual, err := ParseToken(testToken)
	assert.NoError(t, err, "test failed")
	t.Run("ParseToken", func(t *testing.T) {
		assert.Equal(t, user.ID, actual, "ID should be equal")
		assert.NoError(t, err, "should be no errors")
	})
}
