package order_test

import (
	"github.com/0xf0f17a/tau/internal/order"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrderType", func() {
	It("buy order's alternate sell order", func() {
		Expect(order.Buy.Alternate()).To(Equal(order.Sell))
	})
	It("sell order's alternate buy order", func() {
		Expect(order.Sell.Alternate()).To(Equal(order.Buy))
	})
})
