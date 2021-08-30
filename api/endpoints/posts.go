package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdhishaamakhtar/learnGo/api/controllers"
	"github.com/mdhishaamakhtar/learnGo/pkg/models"
)

func healthRoute(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status":  fiber.StatusOK,
		"Message": "Health check successful",
	})
}

func addPost(c *fiber.Ctx) error {
	u := new(models.Posts)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": "Invalid Input Format",
		})
	}
	post, err := controllers.AddPosts(u.Title, u.Description)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Error":   false,
		"Message": "Post Created",
		"Post":    post,
	})
}

func getPost(c *fiber.Ctx) error {
	u := c.Query("id", "")
	if u == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": "Invalid Input Format",
		})
	}
	post, err := controllers.FetchPost(u)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": err.Error(),
		})
	}
	if post.ID == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error":   true,
			"Message": "No such post found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Error":   false,
		"Message": "Post Fetched",
		"Post":    post,
	})
}

func updatePost(c *fiber.Ctx) error {
	u := new(models.Posts)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": "Invalid Input Format",
		})
	}
	if u.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": "No ID provided",
		})
	}
	post, err := controllers.UpdatePost(u.ID, u.Title, u.Description)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Error":   false,
		"Message": "Post Updated",
		"Post":    post,
	})
}

func deletePost(c *fiber.Ctx) error {
	u := c.Query("id", "")
	if u == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": "Invalid Input Format",
		})
	}
	err := controllers.DeletePost(u)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   true,
			"Message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Error":   false,
		"Message": "Post Deleted",
	})
}

func MountRoutes(app *fiber.App) {
	app.Get("/", healthRoute)
	userGroup := app.Group("/api/v1/posts")
	userGroup.Post("/add", addPost)
	userGroup.Get("/fetch", getPost)
	userGroup.Patch("/update", updatePost)
	userGroup.Delete("/delete", deletePost)
}
