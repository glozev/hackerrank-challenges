package main_test

import (
	"bytes"
	"strings"

	. "github.com/glozev/hackerrank-challenges/go/mini_max_sum"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMiniMaxSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MiniMaxSum Suite")
}

var _ = Describe("Finding mini-max sum", func() {
	Context("when reading from input and writing to output", func() {
		It("should work for the sample", func() {
			in := strings.NewReader("1 2 3 4 5")
			out := new(bytes.Buffer)
			Submission(in, out)
			Ω(out.String()).Should(Equal("10 14"))
		})

		It("should work for modified sample", func() {
			in := strings.NewReader("2 4 6 8 10")
			out := new(bytes.Buffer)
			Submission(in, out)
			Ω(out.String()).Should(Equal("20 28"))
		})
	})

	Context("when calculating sum of all integers", func() {
		It("should work with positive", func() {
			Ω(Sum([]int{1, 2, 3})).Should(BeNumerically("==", 6))
		})

		It("should work with large", func() {
			Ω(Sum([]int{1, 2, 1000000000})).Should(BeNumerically("==", 1000000003))
		})
	})

	Context("when calculating mini-max-sum of 4 out of 5 integers", func() {
		It("should work for sample", func() {
			result := Result{10, 14}
			Ω(MiniMaxSum([]int{1, 2, 3, 4, 5})).Should(Equal(result))
		})

		It("should return the same min and max for equal", func() {
			result := Result{4, 4}
			Ω(MiniMaxSum([]int{1, 1, 1, 1, 1})).Should(Equal(result))
		})

		It("should return the same min and max large equal", func() {
			result := Result{4000000000, 4000000000}
			Ω(MiniMaxSum([]int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000})).Should(Equal(result))
		})

		It("should return the same min and max large equal", func() {
			result := Result{1000000003, 2000000002}
			Ω(MiniMaxSum([]int{1, 1, 1, 1000000000, 1000000000})).Should(Equal(result))
		})
	})
})
