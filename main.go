package main

import (
	"goQuiz/main/api"

	"github.com/gin-gonic/gin"

	"goQuiz/main/models"
)

func main() {
	models.QuizInit()
	router := gin.Default()

	router.POST("/create", api.CreateNewQuiz)
	router.GET("/getAll", api.GetAllQuiz)
	router.POST("/check", api.CheckAnswer)

	router.Run(":8000")
}
