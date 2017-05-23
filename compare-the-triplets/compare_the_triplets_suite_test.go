package main_test

import (
	"bytes"
	"strings"

	. "github.com/glozev/hackerrank-challenges/compare-the-triplets"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCompareTheTriplets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CompareTheTriplets Suite")
}

var _ = Describe("when comparing Alice and Bob", func() {
	It("should return equal points", func() {
		alice := Triplet{[3]int{5, 6, 7}}
		bob := Triplet{[3]int{3, 6, 10}}
		Ω(Compare(alice, bob)).Should(Equal("1 1"))
	})

	It("should give more points to Bob", func() {
		alice := Triplet{[3]int{3, 6, 7}}
		bob := Triplet{[3]int{3, 6, 10}}
		Ω(Compare(alice, bob)).Should(Equal("0 1"))
	})

	It("should give more points to Alice", func() {
		alice := Triplet{[3]int{4, 6, 10}}
		bob := Triplet{[3]int{3, 6, 10}}
		Ω(Compare(alice, bob)).Should(Equal("1 0"))
	})

	Context("when reading from input and writing to output", func() {
		It("should work for the sample", func() {
			in := strings.NewReader("5 6 7 3 6 10")
			out := new(bytes.Buffer)
			Submission(in, out)
			Ω(out.String()).Should(Equal("1 1"))
		})
	})
})
