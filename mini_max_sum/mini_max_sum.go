package main

import (
	"fmt"
	"io"
	"os"
)

const InputLength = 5

type Result struct {
	Min uint64
	Max uint64
}

func Submission(stdin io.Reader, stdout io.Writer) {
	input := make([]int, InputLength)

	for i := 0; i < InputLength; i++ {
		fmt.Fscanf(stdin, "%v", &input[i])
	}
	result := MiniMaxSum(input)
	fmt.Fprintf(stdout, "%v %v", result.Min, result.Max)
}

func Sum(array []int) uint64 {
	result := uint64(0)
	for _, value := range array {
		result += uint64(value)
	}
	return result
}

func MiniMaxSum(array []int) Result {
	sum := Sum(array)
	minSum, maxSum := sum, uint64(0)

	for i := 0; i < len(array); i++ {
		partialSum := sum - uint64(array[i])
		if partialSum < minSum {
			minSum = partialSum
		}
		if partialSum > maxSum {
			maxSum = partialSum
		}

	}
	return Result{minSum, maxSum}
}

func main() {
	Submission(os.Stdin, os.Stdout)
}
