package Int32ToIPv4_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInt32ToIPv4(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Int32ToIPv4 Suite")
}
