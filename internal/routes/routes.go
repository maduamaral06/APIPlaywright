package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maduamaral06/APIPlaywright/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.Health)
}