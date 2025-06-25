package controllers

import (
	"api-persona/database"
	"api-persona/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllPersonas(w http.ResponseWriter, r *http.Request) {
	var p []models.Persona
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Persona

	database.DB.First(&persona, id)
	json.NewEncoder(w).Encode(persona)
}
