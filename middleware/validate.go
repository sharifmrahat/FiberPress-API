package middleware

import (
	"fiberpress-api/config"

	"github.com/gofiber/fiber/v2"
)

// ValidateRequest is a generic middleware for struct validation
func ValidateRequest(model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body into the provided model
		if err := c.BodyParser(model); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		// Validate the parsed struct
		errors := config.ValidateStruct(model)
		if len(errors) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
		}

		// Store validated data in Locals so the handler can access it
		c.Locals("validatedData", model)

		// Proceed to the next handler
		return c.Next()
	}
}
