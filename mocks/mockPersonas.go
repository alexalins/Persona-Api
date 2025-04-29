package mocks

import "api-persona/models"

func MockPersonas() []models.Persona {
	var Personas = []models.Persona{
		{
			Id:    1,
			Name:  "Luna",
			Story: "Uma jovem guerreira que luta para proteger seu vilarejo das forças das trevas.",
		},
		{
			Id:    2,
			Name:  "Kai",
			Story: "Um andarilho sábio que busca redenção pelos erros do passado.",
		},
		{
			Id:    3,
			Name:  "Ayla",
			Story: "Uma inventora brilhante que sonha em construir uma máquina do tempo.",
		},
		{
			Id:    4,
			Name:  "Dante",
			Story: "Um caçador de recompensas que vive pelas próprias regras.",
		},
	}

	return Personas
}
