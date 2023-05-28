package TheFeastOfManyBeasts_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTheFeastOfManyBeasts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TheFeastOfManyBeasts Suite")
}
