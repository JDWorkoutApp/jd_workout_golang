package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCorsMiddleware(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

	Cors()(c)

	assert.Equal(t, "*", c.Writer.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "POST, GET, OPTIONS, PUT, DELETE, PATCH", c.Writer.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization", c.Writer.Header().Get("Access-Control-Allow-Headers"))
}

func TestCorsMiddlewareWithOPTIONS(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest(http.MethodOptions, "/", nil)

	Cors()(c)

	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
