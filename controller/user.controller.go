package controller

import "github.com/gofiber/fiber/v2"

func GetUsers(router fiber.Router) {
	router.Get("users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "All users",
		})
	})
}
