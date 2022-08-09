package api

import (
	"goQuiz/main/entities"

	"goQuiz/main/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllQuiz(c *gin.Context) {
	quiz := models.GetAllQuiz()

	c.JSON(http.StatusOK, quiz)
}

func CreateNewQuiz(c *gin.Context) {
	var newQuiz entities.Quiz

	if err := c.BindJSON(&newQuiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		models.CreateQuiz(&newQuiz)
		c.JSON(http.StatusOK, newQuiz)
	}
}
