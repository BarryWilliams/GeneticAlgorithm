package geneticcontroller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGeneticcontroller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Geneticcontroller Suite")
}
