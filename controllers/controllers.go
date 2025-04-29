package controllers

import (
	"api-persona/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func GetAllPersonas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personas)
}

func GetPersonaById(w http.ResponseWriter, r *http.Request) {
	//pegando dados da request
	vars := mux.Vars(r)
	id := vars["id"]

	//procurando dados
	for _, persona := range models.Personas {
		if strconv.Itoa(persona.Id) == id {
			json.NewEncoder(w).Encode(persona)
		}
	}
}
