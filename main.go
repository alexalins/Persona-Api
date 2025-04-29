package main

import (
	"api-persona/mocks"
	"api-persona/models"
	"api-persona/routes"
	"fmt"
)

func main() {
	models.Personas = mocks.MockPersonas()
	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
