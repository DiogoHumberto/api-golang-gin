package database

import (
	"log"

	"github.com/DiogoHumberto/api-go-gin-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	// Connect to your postgres DB.
	stringDeConexao := "host=localhost port=5432 user=docker password=docker dbname=api-erp-escola sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
