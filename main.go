package main

import (
	"api-persona/database"
	"api-persona/routes"
	"fmt"
)

func main() {
	database.InitDatabase()
	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
