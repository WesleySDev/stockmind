package database

import (
	"log"
	"stockmind/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=postgres password=123456 dbname=stonckmind port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar ao banco")
	}

	DB = db

	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Movement{},
	)
	log.Println("Banco conectado!")
}
