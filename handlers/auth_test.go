package handlers

import (
	"blog-platform/database"
	"blog-platform/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	app := Setup()

	// Create a user first
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	user := models.User{
		Username: "testuser",
		Password: string(hashedPassword),
		Role:     "user",
	}
	database.DB.Create(&user)

	app.Post("/login", Login)

	login := models.User{
		Username: "testuser",
		Password: "password",
	}
	body, _ := json.Marshal(login)

	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
