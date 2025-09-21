package main

import (
	"fmt"
	"quiz/quiz"
	"strconv"
)

func ReadQuestion(q *quiz.Question) bool {
	q.Print()

	var input string

	fmt.Printf("Your answer: ")
	fmt.Scanln(&input)

	answer, err := strconv.Atoi(input)

	if err != nil {
		ReadQuestion(q)
	}

	return q.Answer(answer - 1)
}

func main() {
	quiz := quiz.Quiz{
		Questions: []*quiz.Question{
			{
				Text: "Which language is primarily used for web development on the client side?",
				Answers: []*quiz.Answer{
					{Text: "Python", Correct: false},
					{Text: "Go", Correct: false},
					{Text: "C#", Correct: false},
					{Text: "JavaScript", Correct: true},
				},
			},
			{
				Text: "What does HTML stand for?",
				Answers: []*quiz.Answer{
					{Text: "HyperTool Markup Language", Correct: false},
					{Text: "Hyperlink and Text Markup", Correct: false},
					{Text: "HyperText Markup Language", Correct: true},
					{Text: "HighText Machine Language", Correct: false},
				},
			},
			{
				Text: "Which data structure uses LIFO (Last In, First Out)?",
				Answers: []*quiz.Answer{
					{Text: "Queue", Correct: false},
					{Text: "Graph", Correct: false},
					{Text: "Linked List", Correct: false},
					{Text: "Stack", Correct: true},
				},
			},
			{
				Text: "Which symbol is commonly used for single-line comments in most programming languages?",
				Answers: []*quiz.Answer{
					{Text: "<!-- -->", Correct: false},
					{Text: "#!", Correct: false},
					{Text: "//", Correct: true},
					{Text: "/* */", Correct: false},
				},
			},
			{
				Text: "Which programming paradigm focuses on objects and classes?",
				Answers: []*quiz.Answer{
					{Text: "Procedural Programming", Correct: false},
					{Text: "Functional Programming", Correct: false},
					{Text: "Object-Oriented Programming", Correct: true},
					{Text: "Logic Programming", Correct: false},
				},
			},
		},
	}

	correctAnswers := 0

	for _, v := range quiz.Questions {
		if ReadQuestion(v) {
			correctAnswers++
		}
	}

	fmt.Printf("Quiz completed. %d out of %d correct answers.\n", correctAnswers, len(quiz.Questions))
}
