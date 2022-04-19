package main

import (
	"fmt"

	"github.com/arthur-rc18/Api-Go-Rest/database"
	//"github.com/arthur-rc18/Api-Go-Rest/models"
	"github.com/arthur-rc18/Api-Go-Rest/routes"
)

func main() {

	database.DbConnection()

	fmt.Println("Starting the Rest server")

	routes.HandleRequest()
}
