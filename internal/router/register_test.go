package router

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"jd_workout_golang/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	tests.InitDatabase()

	r = gin.Default()
	RegisterUser(r.Group("/api"))
}

func TestRegisterInvalidRequest(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/register", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	response := struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}{}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "register failed", response.Message)
	assert.Equal(t, "missing form body", response.Error)
}

func TestRegisterSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	url := "/api/register"
	username := "go_test_account"
	email := "go_test_account@test.email"
	password := "123456"

	type registerForm struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
		Email    string `json:"email" form:"email" binding:"required"`
	}

	data := registerForm{
		Username: username,
		Password: password,
		Email:    email,
	}

	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 999, w.Code)

	response := struct {
		Message string `json:"message"`
	}{}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "register success", response.Message)
}
