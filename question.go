package main

type Question struct {
	question string
	answer   string
	input    Input
	field    ShortAnswerField
}

func NewQuestion(question string) Question {
	return Question{question: question}
}

func NewShortQuestion(q string) Question {
	question := NewQuestion(q)
	model := NewShortAnswerField()
	question.input = model
	return question
}

func NewLongQuestion(q string) Question {
	question := NewQuestion(q)
	model := NewLongAnswerField()
	question.input = model
	return question
}
