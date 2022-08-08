package api

import (
	"goQuiz/main/entities"

	"encoding/json"

	"io/ioutil"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllQuiz(c *gin.Context) {
	jsonFile, err := ioutil.ReadFile("../data/quizData.json")

	if err == nil {
		var result entities.Quizs

		err2 := json.Unmarshal(jsonFile, &result)

		if err2 == nil {
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
