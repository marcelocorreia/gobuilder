package utils

import (
	"fmt"
)

type Input interface {
	Question(question string) string
	QuestionF(format string, question string) string
}

type Wizard struct  {}

func (w Wizard) Question(question string) string {
	var input string
	fmt.Print(question)
	fmt.Scanln(&input)
	return input
}
func (w Wizard) QuestionF(format string, question string) string {
	var input string
	fmt.Printf(format, question)
	fmt.Scanln(&input)
	return input
}

