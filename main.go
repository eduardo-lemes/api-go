package main

import (
	"fmt"

	"github.com/eduardo-lemes/go-api/database"
	"github.com/eduardo-lemes/go-api/models"
	"github.com/eduardo-lemes/go-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaBancoDados()
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
