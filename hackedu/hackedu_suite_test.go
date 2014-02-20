package hackedu_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHackedu(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hackedu Suite")
}
