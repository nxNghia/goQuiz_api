package api

import (
	"goQuiz/main/entities"

	"goQuiz/main/models"

	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"
)

func GetAllQuiz(c *gin.Context) {
	quiz := models.GetAllQuiz()

	fmt.Println(quiz)

	c.JSON(http.StatusOK, quiz)
}

func CreateNewQuiz(c *gin.Context) {
	var newQuiz entities.Quiz

	if err := c.BindJSON(&newQuiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		models.CreateQuiz(&newQuiz)
		quiz := models.GetAllQuiz()
		c.JSON(http.StatusOK, quiz)
	}
}

func CheckAnswer(c *gin.Context) {
	var answers entities.Answers

	if err := c.BindJSON(&answers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		rightAnswers, total := models.CheckAnswer(answers.Answers)

		c.JSON(http.StatusOK, gin.H{"right": rightAnswers, "total": total})
	}

}
