package value_test

import (
	"github.com/0xf0f17a/tau/internal/value"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Value", func() {
	first := value.NewFloat(34)
	second := value.NewFloat(44)
	Context("Of Float32", func() {
		It("Should have a name method", func() {
			Expect(first.Name()).To(Equal("Float"))
		})
		It("Should have a plus method", func() {
			plus := first.Plus(second)
			result := value.NewFloat(78)
			Expect(plus).To(Equal(result))
		})
		It("Should have a minus method", func() {
			minus := second.Minus(first)
			result := value.NewFloat(10)
			Expect(minus).To(Equal(result))
		})
	})
})
