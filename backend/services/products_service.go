package services

import (
	"stockmind/models"
	"stockmind/repository"
)

func GetProducts() ([]models.Product, error) { // "Busque todos os produtos usando a função GetAllProducts do repository."
	return repository.GetAllProducts() // "Retorne os produtos e o erro, caso haja."
}
