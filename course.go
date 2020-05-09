package main

import (
	"math/rand"
	"time"
)

type Course struct {
	Name        string
	Description string
	Questions   []*Question `json:"questions" yaml:"questions"`
	Scores      []*Score    `json:"submission,omitempty" yaml:"submission,omitempty"`
}

func (c *Course) AddQuestion(q *Question) {
	c.Questions = append(c.Questions, q)
}

func (c *Course) ShuffleQuestions() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c.Questions), func(i, j int) { c.Questions[i], c.Questions[j] = c.Questions[j], c.Questions[i] })
}

func (c *Course) ShuffleAnswers() {
	for _, q := range c.Questions {
		q.ShuffleAnswer()
	}
}

type Question struct {
	Question string
	Type     string    `json:"type" yaml:"type"`
	Answers  []*Answer `json:"answers" yaml:"answers"`
}

func (q *Question) AddAnswer(a *Answer) {
	q.Answers = append(q.Answers, a)
}

func (q *Question) ShuffleAnswer() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q.Answers), func(i, j int) { q.Answers[i], q.Answers[j] = q.Answers[j], q.Answers[i] })
}

type Answer struct {
	Answer   string `json:"answer" yaml:"answer"`
	IsAnswer bool   `json:"is_answer,omitempty" yaml:"is_answer,omitempty"`
}

type Score struct {
	Correct   *Answer `json:"correct" yaml:"correct"`
	Submitted string  `json:"submitted" yaml:"submitted"`
}
