package FibonacciFizzBuzz_test

import (
	. "github.com/andrerademacher/codewars-go/kata/FibonacciFizzBuzz"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("n = 1", func() {
		Expect(FibsFizzBuzz(1)).To(Equal([]string{"1"}))
	})

	It("n = 2", func() {
		Expect(FibsFizzBuzz(2)).To(Equal([]string{"1", "1"}))
	})

	It("n = 3", func() {
		Expect(FibsFizzBuzz(3)).To(Equal([]string{"1", "1", "2"}))
	})
	
	It("n = 4", func() {
		Expect(FibsFizzBuzz(4)).To(Equal([]string{"1", "1", "2", "Fizz"}))
	})

	It("n = 5", func() {
		Expect(FibsFizzBuzz(5)).To(Equal([]string{"1", "1", "2", "Fizz", "Buzz"}))
	})

	It("n = 20", func() {
		Expect(FibsFizzBuzz(20)).To(Equal([]string{"1", "1", "2", "Fizz", "Buzz", "8", "13", "Fizz", "34", "Buzz", "89", "Fizz", "233", "377", "Buzz", "Fizz", "1597", "2584", "4181", "FizzBuzz"}))
	})
})
