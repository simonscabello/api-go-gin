package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/simonscabello/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}
