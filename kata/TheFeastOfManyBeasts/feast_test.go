package TheFeastOfManyBeasts_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/andrerademacher/codewars-go/kata/TheFeastOfManyBeasts"
)

var _ = Describe("Sample Test Cases", func() {
	It("should return the correct value", func() {
		Expect(Feast("great blue heron", "garlic naan")).To(BeTrue(), "Testing 'great blue heron' vs 'garlic naan'")
		Expect(Feast("chickadee", "chocolate cake")).To(BeTrue(), "Testing 'chickadee' vs 'chocolate cake'")
		Expect(Feast("brown bear", "bear claw")).To(BeFalse(), "Testing 'brown bear' vs 'bear claw'")
	})
})
