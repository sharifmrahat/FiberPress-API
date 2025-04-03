package routes

import (
	"fiberpress-api/handlers"
	"fiberpress-api/middleware"
	"fiberpress-api/models"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", middleware.ValidateRequest(&models.Register{}), handlers.RegisterHandler)
	auth.Post("/login", middleware.ValidateRequest(&models.Login{}), handlers.LoginHandler)
	auth.Patch("/makeAdmin/:userId", middleware.AuthMiddleware("admin"), handlers.MakeAdminHandler)

}

