package main

import (
	"fmt"
	"io"
	"os"
)

const Categories = 3

type Triplet struct {
	Scores [Categories]int
}

type Result struct {
	a int
	b int
}

func Compare(a Triplet, b Triplet) string {
	result := Result{}
	for i := 0; i < Categories; i++ {
		if a.Scores[i] > b.Scores[i] {
			result.a++
		} else if a.Scores[i] < b.Scores[i] {
			result.b++
		}
	}
	return fmt.Sprintf("%d %d", result.a, result.b)
}

func Submission(stdin io.Reader, stdout io.Writer) {
	var x, y, z int
	fmt.Fscanf(stdin, "%v %v %v", &x, &y, &z)
	alice := Triplet{[...]int{x, y, z}}
	fmt.Fscanf(stdin, "%v %v %v", &x, &y, &z)
	bob := Triplet{[...]int{x, y, z}}
	fmt.Fprintf(stdout, Compare(alice, bob))
}

func main() {
	Submission(os.Stdin, os.Stdout)
}
