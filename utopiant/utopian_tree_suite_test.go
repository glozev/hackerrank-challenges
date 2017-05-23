package main_test

import (
	. "github.com/glozev/hackerrank-challenges/utopiant"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUtopianTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UtopianTree Suite")
}

var _ = Describe("UtopianTree", func() {
	Describe("Utopian Tree | Find the height after N cycles", func() {
		const InitialHeight int = 1
		Context("With zero growth cycle", func() {
			It("should remain unchanged", func() {
				tree := UtopianTree{InitialHeight, 0}
				Expect(tree.ResultHeight()).To(Equal(InitialHeight))
			})
		})

		Context("after one growth cycle", func() {
			It("should dubles its height", func() {
				tree := UtopianTree{InitialHeight, 1}
				Expect(tree.ResultHeight()).To(Equal(2 * InitialHeight))
			})
		})

		Context("after four growth cycles", func() {
			It("should double its height and grows two meters twice", func() {
				tree := UtopianTree{InitialHeight, 4}
				Expect(tree.ResultHeight()).To(Equal(7))
			})
		})

		Context("after five growth cycles", func() {
			It("should doubles its height three times and grows two meters twice", func() {
				tree := UtopianTree{InitialHeight, 5}
				Expect(tree.ResultHeight()).To(Equal(14))
			})
		})
	})
})
