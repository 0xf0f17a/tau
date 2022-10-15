package order_test

import (
	"github.com/0xf0f17a/tau/internal/instr"
	"github.com/0xf0f17a/tau/internal/order"
	"github.com/0xf0f17a/tau/internal/value"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Order", func() {
	inst := instr.New("InstrumentName")
	price := value.NewFloat(343)
	timeStamp := time.Now()
	var count uint64 = 1
	It("Should create valid buy order", func() {
		ord := order.BuyOrder(inst, price, timeStamp, count)
		Expect(ord.Instr).To(Equal(inst))
		Expect(ord.Price).To(Equal(price))
		Expect(ord.Time).To(Equal(timeStamp))
		Expect(ord.Count).To(Equal(count))
		Expect(ord.Type).To(Equal(order.Buy))
	})
	It("Should create valid sell order", func() {
		ord := order.SellOrder(inst, price, timeStamp, count)
		Expect(ord.Instr).To(Equal(inst))
		Expect(ord.Price).To(Equal(price))
		Expect(ord.Time).To(Equal(timeStamp))
		Expect(ord.Count).To(Equal(count))
		Expect(ord.Type).To(Equal(order.Sell))
	})
	It("Should create valid sell order", func() {
		ord := order.BuyOrder(inst, value.NewFloat(25), timeStamp, uint64(4))
		Expect(ord.Amount()).To(Equal(value.NewFloat(100)))
	})
})
