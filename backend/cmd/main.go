package main

import (
	"stockmind/database"
	"stockmind/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("StockMind API")
	})

	app.Listen(":8080")
}
