package models

import (
	"goQuiz/main/entities"
)

var listQuiz = make([]*entities.Quiz, 0)

func CreateQuiz(quiz *entities.Quiz) bool {
	if quiz.Id != "" && quiz.Answer != 0 && quiz.Content != "" && len(quiz.Option) != 0 {
		if existedQuiz := FindQuiz(quiz.Id); existedQuiz != nil {
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
