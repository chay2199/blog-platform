package handlers

import (
	"blog-platform/database"
	"blog-platform/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	app := Setup()

	app.Post("/posts", CreatePost)

	post := models.Post{
		Title:   "Test Post",
		Content: "This is a test post.",
		Author:  "Test Author",
	}
	body, _ := json.Marshal(post)

	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetPost(t *testing.T) {
	app := Setup()

	// Create a post first
	post := models.Post{
		Title:     "Test Post",
		Content:   "This is a test post.",
		Author:    "Test Author",
		CreatedAt: time.Now(),
	}
	database.DB.Create(&post)

	app.Get("/posts/:id", GetPost)

	req := httptest.NewRequest("GET", "/posts/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetPosts(t *testing.T) {
	app := Setup()

	// Create some posts
	posts := []models.Post{
		{Title: "Post 1", Content: "Content 1", Author: "Author 1", CreatedAt: time.Now()},
		{Title: "Post 2", Content: "Content 2", Author: "Author 2", CreatedAt: time.Now()},
	}
	database.DB.Create(&posts)

	app.Get("/posts", GetPosts)

	req := httptest.NewRequest("GET", "/posts", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdatePost(t *testing.T) {
	app := Setup()

	// Create a post first
	post := models.Post{
		Title:     "Test Post",
		Content:   "This is a test post.",
		Author:    "Test Author",
		CreatedAt: time.Now(),
	}
	database.DB.Create(&post)

	app.Put("/posts/:id", UpdatePost)

	updatedPost := models.Post{
		Title:   "Updated Title",
		Content: "Updated Content",
	}
	body, _ := json.Marshal(updatedPost)

	req := httptest.NewRequest("PUT", "/posts/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeletePost(t *testing.T) {
	app := Setup()

	// Create a post first
	post := models.Post{
		Title:     "Test Post",
		Content:   "This is a test post.",
		Author:    "Test Author",
		CreatedAt: time.Now(),
	}
	database.DB.Create(&post)

	app.Delete("/posts/:id", DeletePost)

	req := httptest.NewRequest("DELETE", "/posts/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
