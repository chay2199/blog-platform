package main

import (
	"blog-platform/database"
	"blog-platform/handlers"
	"blog-platform/middleware"

	_ "blog-platform/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Blogging Platform API
// @version 1.0
// @description This is the API documentation for the Blogging Platform.
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()

	database.InitDatabase()

	// Set up Swagger documentation route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Routes that do not require authentication
	app.Post("/login", handlers.Login)
	app.Post("/users", handlers.CreateUser)

	// Apply the authentication middleware for all routes below
	app.Use(middleware.Authorize())

	// Routes that require a reader or higher role
	app.Get("/posts/:id", middleware.AuthorizeRole("reader", "writer", "admin"), handlers.GetPost)
	app.Get("/posts", middleware.AuthorizeRole("reader", "writer", "admin"), handlers.GetPosts)

	// Routes that require a writer or higher role
	app.Post("/posts", middleware.AuthorizeRole("writer", "admin"), handlers.CreatePost)
	app.Put("/posts/:id", middleware.AuthorizeRole("writer", "admin"), handlers.UpdatePost)
	app.Delete("/posts/:id", middleware.AuthorizeRole("writer", "admin"), handlers.DeletePost)

	// Routes that require admin role
	app.Patch("/users/:id/role", middleware.AuthorizeRole("admin"), handlers.UpdateUserRole)
	app.Delete("/users/:id", middleware.AuthorizeRole("admin"), handlers.DeleteUser)

	app.Listen(":3000")
}
