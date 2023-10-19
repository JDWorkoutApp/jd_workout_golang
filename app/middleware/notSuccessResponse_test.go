package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFailResponseAlertMiddlewareDoingNothing(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	FailResponseAlert()(c)

	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
