package main

import (
	"api_personalidades/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor rest com Go")
	routes.HandleRequest()
}
