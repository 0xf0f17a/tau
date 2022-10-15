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
		It("Should have a multiply method", func() {
			multiply := first.Multiply(second)
			result := value.NewFloat(1496)
			Expect(multiply).To(Equal(result))
		})
		It("Should have a divide method", func() {
			divide := second.Divide(first)
			result := value.NewFloat(1.2941176891326904)
			Expect(divide).To(BeIdenticalTo(result))
		})
		It("Should have a greater than method", func() {
			one := value.NewFloat(32)
			other := value.NewFloat(32)
			another := value.NewFloat(43)
			Expect(another.IsGreaterThan(other)).To(BeTrue())
			Expect(one.IsGreaterThan(another)).To(BeFalse())
		})
		It("Should have a less than method", func() {
			one := value.NewFloat(32)
			other := value.NewFloat(32)
			another := value.NewFloat(43)
			Expect(one.IsLessThan(another)).To(BeTrue())
			Expect(one.IsGreaterThan(other)).To(BeFalse())
		})
		It("Should have a greater than equal to method", func() {
			one := value.NewFloat(32)
			other := value.NewFloat(32)
			another := value.NewFloat(43)
			Expect(one.IsGreaterThanEqualTo(other)).To(BeTrue())
			Expect(one.IsGreaterThanEqualTo(another)).To(BeFalse())
		})
		It("Should have a less than equal to method", func() {
			one := value.NewFloat(32)
			other := value.NewFloat(32)
			another := value.NewFloat(43)
			Expect(one.IsLessThanEqualTo(other)).To(BeTrue())
			Expect(another.IsLessThanEqualTo(one)).To(BeFalse())
		})
		It("Should have a equal to method", func() {
			one := value.NewFloat(32)
			other := value.NewFloat(32)
			another := value.NewFloat(43)
			Expect(one.IsEqualTo(other)).To(BeTrue())
			Expect(another.IsEqualTo(other)).To(BeFalse())
		})
		It("Should have a not equal to method", func() {
			one := value.NewFloat(32)
			other := value.NewFloat(32)
			another := value.NewFloat(43)
			Expect(one.IsNotEqualTo(other)).To(BeFalse())
			Expect(another.IsNotEqualTo(other)).To(BeTrue())
		})
	})
})
