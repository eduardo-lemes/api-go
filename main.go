package main

import (
	"fmt"

	"github.com/eduardo-lemes/go-api/database"
	"github.com/eduardo-lemes/go-api/routes"
)

func main() {

	database.ConectaBancoDados()
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
