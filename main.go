package main

import (
	"fmt"
	"go-fiber-tutorial/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hallo Golang. Long time no see :)")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Service is running!",
		})
	})

	// Controller users
	controller.GetUsers(app.Group("/api/v1/"))

	app.Listen(":8001")
}
