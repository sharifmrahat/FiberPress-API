package main

import (
	"fmt"
	"log"

	"fiberpress-api/config"   // Import config package
	"fiberpress-api/database" // Import database package

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load config
	config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New()

	// Connect to MongoDB
	database.ConnectDB()

	// Define a test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 FiberPress API is Running!")
	})

	// Start the server
	port := ":" + config.AppConfig.ServerPort
	fmt.Println("✅ Server running on http://localhost" + port)
	log.Fatal(app.Listen(port))
}
