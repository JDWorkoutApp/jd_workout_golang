package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateTokenFail(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	c.Request.Header.Set("Authorization", "")

	ValidateToken(c)

	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())
}

func TestValidateResetPasswordFail(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	c.Request.Header.Set("Authorization", "")

	ValidateResetPassword(c)
	assert.Equal(t, http.StatusForbidden, c.Writer.Status())
}
