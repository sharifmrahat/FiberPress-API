package middleware

import (
	"fiberpress-api/config"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware checks for a valid JWT token
func AuthMiddleware(c *fiber.Ctx) error {
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

	// Extract claims and attach user ID to the request
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	// Check token expiration
	if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token expired"})
	}

	// Store user ID in locals for later use
	c.Locals("userId", claims["userId"])

	return c.Next()
}
