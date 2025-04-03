package routes

import (
	"fiberpress-api/handlers"
	"fiberpress-api/middleware"
	"fiberpress-api/models"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/users")

	// Authenticated user routes (admin & author)
	user.Get("/profile", middleware.AuthMiddleware(""), handlers.GetProfile)
	user.Patch("/profile", middleware.AuthMiddleware(""), middleware.ValidateRequest(&models.UpdateProfile{}), handlers.UpdateProfile)
	user.Delete("/profile", middleware.AuthMiddleware(""), handlers.DeleteProfile)

	// Admin-only routes
	user.Get("/", middleware.AuthMiddleware("admin"), handlers.GetAllUsers)
	user.Get("/:id", middleware.AuthMiddleware("admin"), handlers.GetSingleUser)

}
