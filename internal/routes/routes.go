package routes

import (
	"ainyx-user-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Ainyx User API Running",
		})
	})

	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Get("/users", userHandler.GetAllUsers)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
}