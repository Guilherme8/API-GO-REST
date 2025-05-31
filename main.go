package main

import (
	"fmt"

	"github.com/guilherme8/go-api/models"
	"github.com/guilherme8/go-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Dias Vieira", History: "Marechal do Rio de Janeiro"},
		{Id: 2, Name: "Capitão Menezes", History: "Capitão do exército"},
	}

	fmt.Println("Iniciando o servidor com Go")
	routes.HandleRequest()
}
