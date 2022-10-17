package order_test

import (
	"github.com/0xf0f17a/tau/internal/instr"
	"github.com/0xf0f17a/tau/internal/order"
	"github.com/0xf0f17a/tau/internal/value"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	inst      = instr.New("InstrumentName")
	price     = value.NewFloat(343)
	timeStamp = time.Now()
	count     = uint64(1)
)

func TestOrderBuy(t *testing.T) {
	ord := order.BuyOrder(inst, price, timeStamp, count)
	assert.Equal(t, inst, ord.Instr)
	assert.Equal(t, price, ord.Price)
	assert.Equal(t, timeStamp, ord.Time)
	assert.Equal(t, count, ord.Count)
	assert.Equal(t, order.Buy, ord.Type)
}

func TestOrderSell(t *testing.T) {
	ord := order.SellOrder(inst, price, timeStamp, count)
	assert.Equal(t, inst, ord.Instr)
	assert.Equal(t, price, ord.Price)
	assert.Equal(t, timeStamp, ord.Time)
	assert.Equal(t, count, ord.Count)
	assert.Equal(t, order.Sell, ord.Type)
}

func TestOrder_Amount(t *testing.T) {
	p := value.NewFloat(25)
	c := uint64(4)
	amt := value.NewFloat(100)
	ord := order.BuyOrder(inst, p, timeStamp, c)
	assert.Equal(t, amt, ord.Amount())
}
