package middleware

import (
	"bytes"
	"encoding/json"
	"fiberpress-api/config"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

// ValidateRequest ensures the request body matches the expected struct and rejects unknown fields
func ValidateRequest(model interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a new instance of the model (to prevent shared state issues)
		modelInstance := reflect.New(reflect.TypeOf(model).Elem()).Interface()

		decoder := json.NewDecoder(bytes.NewReader(c.Body()))
		decoder.DisallowUnknownFields() // 🚀 This enforces strict schema validation

		if err := decoder.Decode(modelInstance); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload, contains unknown fields"})
		}

		// Validate the parsed struct
		errors := config.ValidateStruct(modelInstance)
		if len(errors) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
		}

		// Store validated data in Locals so the handler can access it
		c.Locals("validatedData", modelInstance)

		// Proceed to the next handler
		return c.Next()
	}
}
