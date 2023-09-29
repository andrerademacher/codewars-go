package EvenOrOdd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEvenOrOdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EvenOrOdd Suite")
}
