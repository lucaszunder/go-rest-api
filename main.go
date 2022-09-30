package main

import (
	"fmt"

	"github.com/lucaszunder/go-rest-api/database"
	"github.com/lucaszunder/go-rest-api/routes"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Starting Server in port 8000")
	routes.HandleRequest()
}
