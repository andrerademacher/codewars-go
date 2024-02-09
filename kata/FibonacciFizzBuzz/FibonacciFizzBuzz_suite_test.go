package FibonacciFizzBuzz_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFibsFizzBuzz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FibonacciFizzBuzz Suite")
}
