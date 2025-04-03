package main

import (
	"fmt"
	"log"

	"fiberpress-api/config"   // Import config package
	"fiberpress-api/database" // Import database package
	"fiberpress-api/middleware"
	"fiberpress-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load config
	config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New()

	// Connect to MongoDB
	database.ConnectDB()

	// Use global middleware
	app.Use(middleware.LoggerMiddleware) // Logging
	// app.Use(middleware.AuthMiddleware)   // Authentication (apply globally if needed)

	// Register Routes
	routes.AuthRoutes(app) // Register authentication routes
	routes.UserRoutes(app) // Register user routes

	// Define a test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 FiberPress API is Running!")
	})

	// Start the server
	port := ":" + config.AppConfig.ServerPort
	fmt.Println("✅ Server running on http://localhost" + port)
	log.Fatal(app.Listen(port))
}
