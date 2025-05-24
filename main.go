package main

import (
	"fmt"

	"github.com/guilherme8/go-api/models"
	"github.com/guilherme8/go-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Dias Vieira", History: "Marechal do Rio de Janeiro"},
	}

	fmt.Println("Iniciando o servidor com Go")
	routes.HandleRequest()
}
