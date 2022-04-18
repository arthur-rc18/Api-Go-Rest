package main

import (
	"fmt"

	"github.com/arthur-rc18/Api-Go-Rest/database"
	"github.com/arthur-rc18/Api-Go-Rest/models"
	"github.com/arthur-rc18/Api-Go-Rest/routes"
)

func main() {

	models.Personalitys = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}

	database.DbConnection()

	fmt.Println("Starting the Rest server")
	routes.HandleRequest()
}
