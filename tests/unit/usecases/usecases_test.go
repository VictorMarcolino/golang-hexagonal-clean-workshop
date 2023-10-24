package usecases_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "USECASE:IN-MEMORY")
}
