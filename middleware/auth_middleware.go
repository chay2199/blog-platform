package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("secret")

// General JWT middleware to check for a valid token
func Authorize() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			log.Println("Authorization header is missing")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Remove "Bearer " prefix if present
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			log.Println("Error parsing token:", err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Check if the token is valid
		if !token.Valid {
			log.Println("Invalid token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Store the token in the context's locals
		c.Locals("user", token)

		// Proceed to the next middleware
		return c.Next()
	}
}

// Middleware to authorize roles
func AuthorizeRole(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the user token from the context
		user := c.Locals("user")
		log.Println(user)
		if user == nil {
			log.Println("User not logged in")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not logged in."})
		}

		// Check if the user is of type *jwt.Token
		token, ok := user.(*jwt.Token)
		if !ok {
			log.Println("Invalid token format")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Check if the token is valid
		if !token.Valid {
			log.Println("Invalid token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Get the claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("Invalid token claims")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Get the user role from the claims
		userRole, ok := claims["role"].(string)
		if !ok {
			log.Println("Role not found in token claims")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Check if the user role is in the required roles
		for _, role := range requiredRoles {
			if userRole == role {
				return c.Next()
			}
		}

		// User role is not in the required roles
		log.Println("Insufficient privileges")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Insufficient privileges"})
	}
}
