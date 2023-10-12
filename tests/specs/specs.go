package specs

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func ExampleSpec() {
	var ()
	BeforeEach(func() {})
	Context("...", func() {
		BeforeEach(func() {})
		It("...", func() {
			Expect(true).To(Equal(true))
		})
	})
}
