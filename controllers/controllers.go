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

func CreatePersona(w http.ResponseWriter, r *http.Request) {
	var newPersona models.Persona
	json.NewDecoder(r.Body).Decode(&newPersona)
	database.DB.Create(&newPersona)
	json.NewEncoder(w).Encode(newPersona)
}

func DeletePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Persona
	database.DB.Delete(&persona, id)
	json.NewEncoder(w).Encode(persona)
}

func UpdatePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Persona
	database.DB.First(&persona, id)
	json.NewDecoder(r.Body).Decode(&persona)
	database.DB.Save(&persona)
	json.NewEncoder(w).Encode(persona)
}
