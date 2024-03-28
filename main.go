package main

import (
	"github.com/simonscabello/api-go-gin/models"
	"github.com/simonscabello/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{
			Nome: "João",
			CPF:  "123.456.789-00",
			RG:   "12.345.678-9",
		},
		{
			Nome: "Maria",
			CPF:  "987.654.321-00",
			RG:   "98.765.432-1",
		},
	}
	routes.HandleRequests()
}
