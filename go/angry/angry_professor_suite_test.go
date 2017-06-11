package main_test

import (
	. "github.com/glozev/hackerrank-challenges/go/angry"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAngry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Angry Suite")
}

var _ = Describe("Angry/AngryProfessor", func() {
	It("should pass dummy test", func() {
		Expect(true).To(Equal(true))
	})

	Describe("AngryProfessor Find if the class get cancelled or not", func() {
		Context("With more than K not attending the course", func() {
			It("should be cancelled", func() {
				prof := AngryProfessor{3, []int{-1, -3, 4, 2}}
				Expect(prof.Result()).To(Equal("YES"))
			})

			It("should be cancelled with most missing", func() {
				prof := AngryProfessor{3, []int{-1, -3, 2, 4, 6, 8, 10, 12, 14}}
				Expect(prof.Result()).To(Equal("YES"))
			})

			It("should be cancelled with all missing", func() {
				prof := AngryProfessor{3, []int{1, 3, 2, 4, 6, 8, 10, 12, 14}}
				Expect(prof.Result()).To(Equal("YES"))
			})

		})

		Context("With less than K attending the course", func() {
			It("should start", func() {
				prof := AngryProfessor{2, []int{0, -1, 2, 1}}
				Expect(prof.Result()).To(Equal("NO"))
			})

			It("should start with all just on time ", func() {
				prof := AngryProfessor{3, []int{0, 0, 0, 0}}
				Expect(prof.Result()).To(Equal("NO"))
			})

			It("should start with all earlier ", func() {
				prof := AngryProfessor{3, []int{-1, -2, -3, -4}}
				Expect(prof.Result()).To(Equal("NO"))
			})

			It("should start without zero threshold ", func() {
				prof := AngryProfessor{0, []int{-1, -2, -3, -4}}
				Expect(prof.Result()).To(Equal("NO"))
			})
		})
	})
})
