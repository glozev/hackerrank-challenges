package main_test

import (
	"bytes"
	"strings"

	. "github.com/glozev/hackerrank-challenges/birthday_cake_candles"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBirthdayCakeCandles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BirthdayCakeCandles Suite")
}

var _ = Describe("BirthdayCakeCandles", func() {
	Context("when reading from input and writing to output", func() {
		It("should work for the sample", func() {
			in := strings.NewReader("CHANE_ME_INPUT")
			out := new(bytes.Buffer)
			Submission(in, out)
			Ω(out.String()).Should(Equal("CHANGE_ME_OUTPUT"))
		})
	})
})
