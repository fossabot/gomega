package matchers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/matchers"
)

var _ = Describe("HaveLen", func() {
	Context("when passed a supported type", func() {
		It("should do the right thing", func() {
			Ω("").Should(HaveLen(0))
			Ω("AA").Should(HaveLen(2))

			Ω([0]int{}).Should(HaveLen(0))
			Ω([2]int{1, 2}).Should(HaveLen(2))

			Ω([]int{}).Should(HaveLen(0))
			Ω([]int{1, 2, 3}).Should(HaveLen(3))

			Ω(map[string]int{}).Should(HaveLen(0))
			Ω(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}).Should(HaveLen(4))

			c := make(chan bool, 3)
			Ω(c).Should(HaveLen(0))
			c <- true
			c <- true
			Ω(c).Should(HaveLen(2))
		})
	})

	Context("when passed a typed nil", func() {
		It("should error", func() {
			var nilMap map[string]int
			success, _, err := (&HaveLenMatcher{Count: 0}).Match(nilMap)
			Ω(success).Should(BeFalse())
			Ω(err).Should(HaveOccured())

			var nilSlice []int
			success, _, err = (&HaveLenMatcher{Count: 0}).Match(nilSlice)
			Ω(success).Should(BeFalse())
			Ω(err).Should(HaveOccured())
		})
	})

	Context("when passed an unsupported type", func() {
		It("should error", func() {
			success, _, err := (&HaveLenMatcher{Count: 0}).Match(0)
			Ω(success).Should(BeFalse())
			Ω(err).Should(HaveOccured())

			success, _, err = (&HaveLenMatcher{Count: 0}).Match(nil)
			Ω(success).Should(BeFalse())
			Ω(err).Should(HaveOccured())
		})
	})
})