package models

import (
	"fmt"

	"goQuiz/main/entities"

	"encoding/json"

	"io/ioutil"
)

var listQuiz = make([]*entities.Quiz, 0)

func CreateQuiz(quiz *entities.Quiz) bool {
	if quiz.Id != "" && quiz.Answer != 0 && quiz.Content != "" && len(quiz.Option) != 0 {
		if existedQuiz := FindQuiz(quiz.Id); existedQuiz == nil {
			listQuiz = append(listQuiz, quiz)
			return true
		}
	}

	return false
}

func FindQuiz(id string) *entities.Quiz {
	for _, quiz := range listQuiz {
		if quiz.Id == id {
			return quiz
		}
	}

	return nil
}

func GetAllQuiz() []*entities.Quiz {
	return listQuiz
}

func QuizInit() bool {
	fmt.Println("init")
	jsonFile, err := ioutil.ReadFile("./data/quizData.json")

	if err == nil {
		var result entities.Quizs

		err2 := json.Unmarshal(jsonFile, &result)

		if err2 == nil {
			for i := 0; i < len(result.Quizs); i++ {
				listQuiz = append(listQuiz, &result.Quizs[i])
			}

			return true
		} else {
			fmt.Println(err2.Error())
			return false
		}
	} else {
		fmt.Println(err.Error())
	}

	return false
}

func CheckAnswer(answers []int) (int, int) {
	rightAnswers := 0

	for index, answer := range answers {
		if answer == listQuiz[index].Answer {
			rightAnswers++
		}
	}

	return rightAnswers, len(answers)
}

func UpdateQuiz(newQuiz entities.Quiz) entities.Quiz {
	return newQuiz
}

func RemoveQuiz(id string) bool {
	// do something
	return true
}
