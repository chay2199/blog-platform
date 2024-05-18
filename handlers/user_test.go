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

func TestCreateUser(t *testing.T) {
	app := Setup()

	app.Post("/users", CreateUser)

	user := models.User{
		Username: "testuser",
		Password: "password",
	}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateUserRole(t *testing.T) {
	app := Setup()

	// Create a user first
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	user := models.User{
		Username: "testuser",
		Password: string(hashedPassword),
		Role:     "user",
	}
	database.DB.Create(&user)

	app.Patch("/users/:id/role", UpdateUserRole)

	req := httptest.NewRequest("PATCH", "/users/1/role?role=admin", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteUser(t *testing.T) {
	app := Setup()

	// Create a user first
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	user := models.User{
		Username: "testuser",
		Password: string(hashedPassword),
	}
	database.DB.Create(&user)

	app.Delete("/users/:id", DeleteUser)

	req := httptest.NewRequest("DELETE", "/users/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
