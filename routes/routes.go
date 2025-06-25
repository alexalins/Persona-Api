package routes

import (
	"api-persona/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/personas", controllers.GetAllPersonas).Methods("Get")
	r.HandleFunc("/persona/{id}", controllers.GetPersonaById).Methods("Get")
	r.HandleFunc("/persona", controllers.CreatePersona).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
