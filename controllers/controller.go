package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/simonscabello/api-go-gin/models"
)

func ExibeAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"mensagem": "Olá, " + nome,
	})
}
