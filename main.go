package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error opening %s\n", *csvFilename))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file")
	}
	Quiz := parseProblems(lines)

	var counter int
	var answer string
	for i := 0; i < Quiz.total; i++ {
		fmt.Printf(
			"Problem number %d out of %d.\n%s = ",
			i+1,
			len(Quiz.problems),
			Quiz.getQuestion(i),
		)

		fmt.Scanf("%s\n", &answer)
		if Quiz.checkAnswer(i, answer) {
			counter++
		}
	}

	fmt.Printf("Quiz finished!\nYou scored %d out of %d\n", counter, Quiz.total)
}

type problem struct {
	question string
	answer   string
}

type quiz struct {
	problems []problem
	total    int
}

func (q *quiz) addProblem(p problem) {
	q.problems = append(q.problems, p)
	q.total++
}

func (q *quiz) getQuestion(n int) string {
	return q.problems[n].question
}

func (q *quiz) getAnswer(n int) string {
	return q.problems[n].answer
}

func (q *quiz) checkAnswer(n int, a string) bool {
	return q.problems[n].answer == a
}

func parseProblems(lines [][]string) quiz {
	var res quiz
	for _, line := range lines {
		res.addProblem(problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		})
	}
	return res
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
