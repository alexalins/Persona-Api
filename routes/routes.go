package routes

import (
	"api-persona/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/persona", controllers.GetAllPersonas)
	log.Fatal(http.ListenAndServe(":8000", r))
}
