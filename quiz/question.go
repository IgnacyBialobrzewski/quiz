package quiz

import "fmt"

type Question struct {
	Text    string
	Answers []*Answer
}

func (q *Question) Answer(idx int) bool {
	if idx >= len(q.Answers) || idx < 0 {
		return false
	}

	return q.Answers[idx].Correct
}

func (q *Question) Print() {
	fmt.Println(q.Text)

	for i, v := range q.Answers {
		fmt.Printf("%d. %s\n", i+1, v.Text)
	}
}
