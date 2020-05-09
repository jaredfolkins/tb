package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/manifoldco/promptui"
)

const example = `
name: Oregon Trail
description: A 10 question quiz on the Oregon Trail
questions:

- question: 'Around how long is the Oregon Trail?'
  type: multiple-choice
  answers:
  - answer: 100 miles 
  - answer: 500 miles
  - answer: 1,000 miles
  - answer: 2,000 miles
    is_answer: true
  - answer: 4,000 miles

- question: 'In what state did the Oregon Trail begin?'
  type: multiple-choice
  answers:
  - answer: New York
  - answer: Missouri
    is_answer: true
  - answer: Florida
  - answer: California
  - answer: Wyoming

- question: 'What was the main vehicle used to carry belongings by pioneers on the Oregon Trail?'
  type: multiple-choice
  answers:
  - answer: Train
  - answer: Steamboat
  - answer: Covered Wagon
    is_answer: true
  - answer: Jeep Cherokee
  - answer: Canoe

- question: 'True or False: The main danger to pioneers on the trail was Native Americans.'
  type: true-false
  answers:
  - answer: "True"
  - answer: "False"
    is_answer: true

- question: 'Around how long did it typically take for a wagon train to travel the Oregon Trail?'
  type: multiple-choice
  answers:
  - answer: 2 weeks
  - answer: 1 month
  - answer: 5 months
  - answer: 1 year
    is_answer: true
  - answer: 2 year

- question: 'Which of the following states did the Oregon Trail NOT pass through?'
  type: multiple-choice
  answers:
  - answer: California
    is_answer: true
  - answer: Nebraska
  - answer: Wyoming
  - answer: Idaho
  - answer: Oregon

- question: 'During what century was the Oregon Trail most traveled?'
  type: multiple-choice
  answers:
  - answer: 1600s
  - answer: 1700s
  - answer: 1800s
    is_answer: true
  - answer: 1900s
  - answer: 2000s

- question: 'In what state did the Oregon Trail end?'
  type: multiple-choice
  answers:
  - answer: California
  - answer: Washington
  - answer: Arizona
  - answer: New Mexico
  - answer: Oregon
    is_answer: true

- question: 'What was the main cause of death to pioneers on the trail?'
  type: multiple-choice
  answers:
  - answer: Starvation
  - answer: Disease
    is_answer: true
  - answer: Native American Attacks
  - answer: Tornadoes
  - answer: Fire

- question: 'What was the main item that pioneers brought with them in their covered wagons?'
  type: multiple-choice
  answers:
  - answer: Food
    is_answer: true
  - answer: Furniture
  - answer: Clothing
  - answer: Books
  - answer: Electronics
`
const helpTxt = `TestBuilder %s
A simple binary that allows easy construction of tests using yaml.

Example
%s

`

const empty = ""

var c *Course

func main() {
	// flags
	help := flag.Bool("help", false, "Print the help for TestBuilder")
	example := flag.Bool("example", false, "Create an example.yml in the current directory")
	file := flag.String("file", "", "Take a test using the specified file (/some/path/to/file/example.yml)")
	flag.Parse()

	if *help {
		displayHelp()
		return
	}

	if *example {
		if err := createExampleFile(); err != nil {
			fmt.Printf("Error creating example.yml file {%v}\n", err)
		}
		return
	}

	if fileExists(*file) {
		runTest(*file)
		for _, score := range c.Scores {
			response := "incorrect!"
			if score.Correct.Answer == score.Submitted {
				response = "correct!"
			}
			fmt.Printf("You submitted {%s} and the answer is {%s} which was %s\n", score.Submitted, score.Correct.Answer, response)
		}

		return
	}

	flag.PrintDefaults()
	return
}

func runTest(filepath string) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error creating example.yml file {%v}\n", err)
		return
	}
	err = yaml.Unmarshal(f, &c)
	c.ShuffleQuestions()
	c.ShuffleAnswers()

	for _, q := range c.Questions {
		var answers []string
		score := &Score{}
		for _, a := range q.Answers {
			answers = append(answers, a.Answer)
		}

		prompt := promptui.Select{
			HideSelected: true,
			Label:        q.Question,
			Items:        answers,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		for _, a := range q.Answers {
			if a.IsAnswer == true {
				score.Correct = a
				score.Submitted = result
				c.Scores = append(c.Scores, score)
			}
		}

		fmt.Printf("%s %s\n", q.Question, result)
	}

}

func createExampleFile() error {
	c := exampleCourse()
	y, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("example.yml", []byte(y), 0644)
}

func exampleCourse() *Course {
	c := &Course{}
	err := yaml.Unmarshal([]byte(example), &c)
	if err != nil {
		log.Fatalln(err)
	}
	return c
}

func exampleBuildCourse() *Course {
	c := &Course{
		Name:        "HTML v1",
		Description: "This Course is questions about HTML",
	}

	q1 := &Question{
		Type:     "multiple-choice",
		Question: "What color does the hexcode #000000 represent in HTML?",
	}

	q1.AddAnswer(&Answer{Answer: "Red"})
	q1.AddAnswer(&Answer{Answer: "White"})
	q1.AddAnswer(&Answer{Answer: "Green"})
	q1.AddAnswer(&Answer{Answer: "Black", IsAnswer: true})
	q1.ShuffleAnswer()

	q2 := &Question{
		Type:     "true-false",
		Question: "The <html> tag is the outer most tag in HTML5.",
	}

	q2.AddAnswer(&Answer{Answer: "True", IsAnswer: true})
	q2.AddAnswer(&Answer{Answer: "False"})

	c.AddQuestion(q1)
	c.AddQuestion(q2)
	c.ShuffleQuestions()
	return c
}

func displayHelp() {
	c := exampleCourse()
	y, _ := yaml.Marshal(&c)
	fmt.Printf(helpTxt, "v0.0.0", y)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
