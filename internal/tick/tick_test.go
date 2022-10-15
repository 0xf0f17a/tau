package tick_test

import (
	"github.com/0xf0f17a/tau/internal/instr"
	"github.com/0xf0f17a/tau/internal/tick"
	"github.com/0xf0f17a/tau/internal/value"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Tick", func() {
	inst := "InstrumentName"
	var price float32 = 343.3
	timeStamp := time.Now()
	var count uint64 = 1
	t := tick.New(instr.New(inst), value.NewFloat(price), timeStamp, count)
	It("Should return all valid values", func() {
		Expect(t.Instr.String()).To(Equal("InstrumentName"))
		Expect(t.Price).To(Equal(value.NewFloat(price)))
		Expect(t.Time).To(Equal(timeStamp))
		Expect(t.Count).To(Equal(count))
	})
})
