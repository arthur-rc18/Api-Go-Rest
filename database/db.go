package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnection() {

	connectionString := "host=localhost user=root password=root dbname=root port=54321 sslmode=disable "
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Connection to database failed")
	}
}
