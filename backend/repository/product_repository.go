package repository

import (
	"stockmind/database"
	"stockmind/models"
)

func GetAllProducts() ([]models.Product, error) {

	var products []models.Product

	err := database.DB.Find(&products).Error //"Busque todos os produtos da tabela products."

	return products, err
}
