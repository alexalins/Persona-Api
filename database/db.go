package database

import (
	"api-persona/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	connectValue := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(connectValue))
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	DB.AutoMigrate(&models.Persona{})
}
