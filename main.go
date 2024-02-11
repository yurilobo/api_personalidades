package main

import (
	"api_personalidades/models"
	"api_personalidades/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
		{Id: 3, Nome: "Nome 3", Historia: "Historia 3"},
	}

	fmt.Println("Iniciando o servidor rest com Go")
	routes.HandleRequest()
}
