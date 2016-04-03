package utils

import (
	"fmt"
)
/*
	Returns the input of the question
*/
func Question(question string) string {
	var input string
	fmt.Print(question)
	fmt.Scanln(&input)
	return input
}
func QuestionF(format string, question string) string {
	var input string
	fmt.Printf(format, question)
	fmt.Scanln(&input)
	return input
}