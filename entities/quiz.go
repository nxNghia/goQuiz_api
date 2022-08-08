package entities

import (
	"fmt"
)

type Quiz struct {
	Id      string   `json:"id"`
	Content string   `json:"content"`
	Answer  int      `json:"answer"`
	Option  []string `json:"option"`
}

func (quiz Quiz) ToString() string {
	return fmt.Sprintf("id: %s\ncontent: %s\nanswer: %s", quiz.Id, quiz.Content, quiz.Option[quiz.Answer])
}
