package main

func main() {

}

type problem struct {
	question string
	answer   string
}

type quiz struct {
	problems []problem
	total    int
	current  int
}

func (q quiz) addProblem(p problem) {
	q.problems = append(q.problems, p)
	q.total++
	q.current++
}
