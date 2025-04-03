package middleware

import (
	"fiberpress-api/config"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware checks for a valid JWT token and verifies the user role
func AuthMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		// Check if the token is provided
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		// Extract token (Bearer <token>)
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		// Parse the JWT token
		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		// Check token expiration
		if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token expired"})
		}

		// Extract user role, default to "author"
		role, ok := claims["role"].(string)
		
		if !ok {
			role = "author"
		}

		// Store user ID and role in locals for later use
		c.Locals("userId", claims["userId"])
		c.Locals("role", role)



		// Check if the user has the required role
		if requiredRole != "" && role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not authorized to perform this action"})
		}

		return c.Next()
	}
}
