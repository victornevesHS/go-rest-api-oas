package main

import (
	"fmt"

	"github.com/victornevesHS/go-rest-api-oas/models"
	"github.com/victornevesHS/go-rest-api-oas/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Albert Einstein", Historia: "Físico alemão"},
		{Id: 2, Nome: "Isaac Newton", Historia: "Físico e matemático inglês"},
	}

	fmt.Println("iniciando o servidor...")
	routes.HnadleRequests()
}
