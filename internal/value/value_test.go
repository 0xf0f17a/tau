package value_test

import (
	"github.com/0xf0f17a/tau/internal/value"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Value", func() {
	f32Value := value.NewFloat(34)
	another := value.NewFloat(44)
	Context("Of Float32", func() {
		It("Should have a name method", func() {
			Expect(f32Value.Name()).To(Equal("Float"))
		})
		It("Should have a plus method", func() {
			plus := f32Value.Plus(another)
			result := value.NewFloat(78)
			Expect(plus).To(Equal(result))
			Expect(&plus).NotTo(Equal(f32Value))
		})
	})
})
