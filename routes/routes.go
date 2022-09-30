package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucaszunder/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.ListPersonalities).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.GetOnePersonality).Methods("Get")
	r.HandleFunc("/personalidades", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controllers.UpdatePersonality).Methods("Patch")
	r.HandleFunc("/personalidades/{id}", controllers.DeletePersonality).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}
