package routes

import (
	"api-persona/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/persona", controllers.GetAllPersonas)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
