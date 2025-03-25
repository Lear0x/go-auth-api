package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Lear0x/go-auth-api/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	os.Setenv("APP_ENV", "test")
	config.LoadEnv()
	config.ConnectDB()
	config.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")

	r := gin.Default()

	api := r.Group("/19ebe88a-e0ce-42bc-8dcf-d5206d0658ad")
	{
		api.POST("/register", Register)
		api.POST("/login", Login)
	}

	return r
}

func TestRegisterAndLogin(t *testing.T) {
	router := setupTestRouter()

	// ---------- Test /register ----------
	registerBody := map[string]string{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "securepass",
	}
	bodyJSON, _ := json.Marshal(registerBody)

	req := httptest.NewRequest("POST", "/19ebe88a-e0ce-42bc-8dcf-d5206d0658ad/register", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)

	// ---------- Test /login ----------
	loginBody := map[string]string{
		"email":    "test@example.com",
		"password": "securepass",
	}
	bodyJSON, _ = json.Marshal(loginBody)

	req = httptest.NewRequest("POST", "/19ebe88a-e0ce-42bc-8dcf-d5206d0658ad/login", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	var respBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &respBody)

	_, tokenExists := respBody["token"]
	assert.True(t, tokenExists, "la réponse doit contenir un token")
}
