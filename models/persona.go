package models

type Persona struct {
	Id    int    `json:"id"`
	Name  string `json:"nome"`
	Story string `json:"historia"`
}

var Personas []Persona
