package handlers

import (
	"blog-platform/database"
	"blog-platform/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreatePost creates a new blog post
// @Summary Create a new post
// @Description Create a new blog post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body models.Post true "Blog post"
// @Success 201 {object} models.Post
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /posts [post]
func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	post.CreatedAt = time.Now()
	if err := database.DB.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create post"})
	}
	return c.JSON(post)
}

// GetPost retrieves a blog post by ID
// @Summary Get a post
// @Description Get a blog post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /posts/{id} [get]
func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot get post"})
	}
	return c.JSON(post)
}

// GetPosts retrieves a list of blog posts
// @Summary Get posts
// @Description Get a list of blog posts
// @Tags posts
// @Accept json
// @Produce json
// @Param author query string false "Author"
// @Param date query string false "Date (YYYY-MM-DD)"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {array} models.Post
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /posts [get]
func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	query := database.DB

	author := c.Query("author")
	if author != "" {
		query = query.Where("author = ?", author)
	}

	date := c.Query("date")
	if date != "" {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err == nil {
			query = query.Where("date(created_at) = ?", parsedDate.Format("2006-01-02"))
		}
	}

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	query.Offset(offset).Limit(limit).Find(&posts)
	return c.JSON(posts)
}

// UpdatePost updates an existing blog post
// @Summary Update a post
// @Description Update an existing blog post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body models.Post true "Blog post"
// @Success 200 {object} models.Post
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /posts/{id} [put]
func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot get post"})
	}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := database.DB.Save(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update post"})
	}
	return c.JSON(post)
}

// DeletePost deletes a blog post by ID
// @Summary Delete a post
// @Description Delete a blog post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 204 {object} nil
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /posts/{id} [delete]
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Post{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete post"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
