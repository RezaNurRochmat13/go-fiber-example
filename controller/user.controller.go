package controller

import "github.com/gofiber/fiber/v2"

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "All users",
	})
}

func GetSingleUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Single users",
	})
}
