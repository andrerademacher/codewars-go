package EvenOrOdd_test

import (
	. "github.com/andrerademacher/codewars-go/kata/EvenOrOdd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Even", func() {
	It("should return the correct value", func() {
		Expect(EvenOrOdd(0)).To(Equal("Even"))
		Expect(EvenOrOdd(1)).To(Equal("Odd"))
		Expect(EvenOrOdd(2)).To(Equal("Even"))
		Expect(EvenOrOdd(-1)).To(Equal("Odd"))
		Expect(EvenOrOdd(-2)).To(Equal("Even"))
	})
})
