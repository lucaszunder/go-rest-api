package main

import (
	"fmt"

	"github.com/lucaszunder/go-rest-api/routes"
)

func main() {
	fmt.Println("Starting Server in port 8000")
	routes.HandleRequest()
}
