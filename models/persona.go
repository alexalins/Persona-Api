package models

type Persona struct {
	Name  string `json:"nome"`
	Story string `json:"historia"`
}

var Personas []Persona
