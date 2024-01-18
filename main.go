package main

import (
	"api_personalidades/models"
	"api_personalidades/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 1", Historia: "Historia 1"},
	}

	fmt.Println("Iniciando o servidor rest com Go")
	routes.HandleRequest()
}
