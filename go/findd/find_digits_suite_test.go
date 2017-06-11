package main_test

import (
	. "github.com/glozev/hackerrank-challenges/go/findd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFindd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FindDigits Suite")
}

var _ = Describe("Find Digits", func() {
	Describe("FindDigits | Find the count of digits in N that exactly divide it", func() {
		Context("without repeated digits", func() {
			It("should return the number of exact dividers", func() {
				dig := FindDigits{12}
				Expect(dig.Count()).To(Equal(2))
			})
		})

		Context("with repeated digits", func() {
			It("should return the repeated and all other exact dividers", func() {
				dig := FindDigits{1012}
				Expect(dig.Count()).To(Equal(3))
			})
		})
	})
})
