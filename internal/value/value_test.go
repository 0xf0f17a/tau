package value_test

import (
	"github.com/0xf0f17a/tau/internal/value"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Value", func() {
	f32Value := value.NewFloat(34)
	Context("Of Float32", func() {
		It("It Should have a name method", func() {
			Expect(f32Value.Name()).To(Equal("Float"))
		})
	})
})
