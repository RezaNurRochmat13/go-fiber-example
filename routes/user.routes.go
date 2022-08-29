package routes

import (
	"go-fiber-tutorial/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	router.Get("users", controller.GetUsers)
	router.Get("users/:id", controller.GetSingleUsers)
	router.Post("users", controller.CreateNewUsers)
	router.Put("users/:id", controller.UpdateUsers)
	router.Delete("users/:id", controller.DeleteUsers)
}
