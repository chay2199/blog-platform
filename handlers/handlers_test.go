package handlers

import (
	"blog-platform/database"
	"blog-platform/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup initializes a new Fiber app and database for testing
func Setup() *fiber.App {
	// Initialize the test database
	database.DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.DB.AutoMigrate(&models.Post{}, &models.User{})

	// Initialize the Fiber app
	app := fiber.New()

	return app
}
