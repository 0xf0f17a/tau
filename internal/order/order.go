package order

import (
	"github.com/0xf0f17a/tau/internal/instr"
	"github.com/0xf0f17a/tau/internal/value"
	"time"
)

type Order struct {
	Instr instr.Instrument
	Price value.Value
	Time  time.Time
	Count uint64
	Type  Type
}

func BuyOrder(instr instr.Instrument, price value.Value, time time.Time, count uint64) *Order {
	return newOrder(instr, price, time, count, Buy)
}

func SellOrder(instr instr.Instrument, price value.Value, time time.Time, count uint64) *Order {
	return newOrder(instr, price, time, count, Sell)
}

func newOrder(instr instr.Instrument, price value.Value, time time.Time, count uint64, t Type) *Order {
	return &Order{instr, price, time, count, t}
}

func (o *Order) Amount() value.Value {
	return o.Price.Multiply(value.NewFloat(float32(o.Count)))
}
