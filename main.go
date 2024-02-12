package main

import (
	"api_personalidades/database"
	"api_personalidades/models"
	"api_personalidades/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor rest com Go")
	routes.HandleRequest()
}

//aqui nesse commit eu tive alguns problemas,eu tinha um postgress rodando local, apontando para porta 5432, tive que da um stop
//iniciei o composer na porta 5435 para evitar o conflito, e deu erro
//depois matei o postgres local, mudei as portadas do docker e rodou
