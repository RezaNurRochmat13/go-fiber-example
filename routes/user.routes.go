package routes

import (
	"go-fiber-tutorial/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	router.Get("users", controller.GetUsers)
}
