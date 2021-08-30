package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mdhishaamakhtar/learnGo/api/endpoints"
	"log"
)

func main() {
	app := fiber.New()

	// Initializing any middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	// Mounting routes
	endpoints.MountRoutes(app)

	// Running the server
	log.Fatal(app.Listen(":3000"))
}
