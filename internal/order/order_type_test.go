package order_test

import (
	"github.com/0xf0f17a/tau/internal/order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderType_Alternate(t *testing.T) {
	assert.Equal(t, order.Sell, order.Buy.Alternate())
	assert.Equal(t, order.Buy, order.Sell.Alternate())
}
