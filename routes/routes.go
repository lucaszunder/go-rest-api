package routes

import (
	"log"
	"net/http"

	"github.com/lucaszunder/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
