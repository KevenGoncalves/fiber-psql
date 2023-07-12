package users

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App, controller *UserController) {
	router := app.Group("/user")

	router.Post("/", controller.Create)
	router.Get("/", controller.List)
	router.Get("/:id", controller.Get)
	router.Put("/:id", controller.Update)
	router.Delete("/:id", controller.Delete)
}
