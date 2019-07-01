package main

import (
	"fmt"
	"testing"
)

func TestAddProblem(t *testing.T) {
	var q quiz
	q.addProblem(problem{
		question: "hello",
		answer:   "word",
	})
	fmt.Printf("Number of quizes: %d\n", q.total)
	if q.problems[0].question != "hello" {
		t.Error(q)
	}
}
