package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/maduamaral06/APIPlaywright/internal/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)
	log.Println("Server is running on http://localhost:3000")

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
	
