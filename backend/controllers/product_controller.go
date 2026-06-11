package controllers

import (
	"stockmind/services"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {

	products, err := services.GetProducts()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Erro ao buscar produtos",
		})
	}
	return c.JSON(products)
}
