package main

import (
	"fmt"

	"github.com/arthur-rc18/Api-Go-Rest/database"
	"github.com/arthur-rc18/Api-Go-Rest/routes"
)

func main() {

	// Connecing to the DB
	database.DbConnection()

	fmt.Println("Starting the Rest server")

	// Hnadling the routes
	routes.HandleRequest()
}
