package endpoints

import "github.com/gofiber/fiber/v2"

func healthRoute(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status":  fiber.StatusOK,
		"Message": "Health check successful",
	})
}

func MountRoutes(app *fiber.App) {
	app.Get("/", healthRoute)
}
