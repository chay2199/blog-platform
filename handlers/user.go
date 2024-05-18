package handlers

import (
	"blog-platform/database"
	"blog-platform/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot hash password"})
	}
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create user"})
	}

	return c.JSON(user)
}

// UpdateUserRole updates a user's role
// @Summary Update user role
// @Description Update a user's role
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param role query string true "New role"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /users/{id}/role [patch]
func UpdateUserRole(c *fiber.Ctx) error {
	id := c.Params("id")
	role := c.Query("role")
	if role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Role is required"})
	}

	var user models.User
	if result := database.DB.First(&user, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot get user"})
	}

	user.Role = role
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update user role"})
	}

	return c.JSON(user)
}

// DeleteUser deletes a user by ID
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {object} nil
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if result := database.DB.Delete(&models.User{}, id); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete user"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
