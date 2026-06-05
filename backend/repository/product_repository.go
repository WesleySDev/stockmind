package repository

import (
	"stockmind/database"
	"stockmind/models"
)

func GetAllProducts() ([]models.Product, error) { // "Declare uma variável para armazenar os produtos."

	var products []models.Product // "Busque todos os produtos usando a função Find do GORM, passando a variável products como referência."

	err := database.DB.Find(&products).Error // "Retorne os produtos e o erro, caso haja."

	return products, err
}
