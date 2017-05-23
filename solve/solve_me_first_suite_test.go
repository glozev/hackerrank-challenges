package main_test

import (
	. "github.com/glozev/hackerrank-challenges/solve"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSolveMeFirst(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SolveMeFirst Suite")
}

var _ = Describe("SolveMeFirst", func() {
	It("should sum two positive integers", func() {
		Ω(SolveMeFirst(1, 2)).Should(BeNumerically("==", 3))
	})
})
