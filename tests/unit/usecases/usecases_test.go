package usecases_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "USECASE:IN-MEMORY")
}
