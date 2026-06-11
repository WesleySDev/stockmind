package routes

import (
	"stockmind/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/products", controllers.GetProducts)
}
