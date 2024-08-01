package main

import (
	"fmt"

	"github.com/victornevesHS/go-rest-api-oas/models"
	"github.com/victornevesHS/go-rest-api-oas/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Albert Einstein", Historia: "Físico alemão"},
		{Nome: "Isaac Newton", Historia: "Físico e matemático inglês"},
	}

	fmt.Println("iniciando o servidor...")
	routes.HnadleRequests()
}
