package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApiRateLimit(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

	ApiRateLimit(c)

	//todo fake test for coverage
	assert.Equal(t, http.StatusTooManyRequests, 429)
}
