package routes

import (
	"api-persona/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/personas", controllers.GetAllPersonas)
	r.HandleFunc("/persona/{id}", controllers.GetPersonaById)
	log.Fatal(http.ListenAndServe(":8000", r))
}
