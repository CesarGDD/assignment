package routes

import (
	"cesargdd/rest-api-assigment/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/")
	api.Get("registrations", controllers.GetRegistrations)
	api.Post("registration", controllers.CreateRegistrations)
}
