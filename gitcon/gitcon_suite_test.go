package gitcon_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGitcon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gitcon Suite")
}
