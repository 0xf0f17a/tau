package tick_test

import (
	"github.com/0xf0f17a/tau/internal/instr"
	"github.com/0xf0f17a/tau/internal/tick"
	"github.com/0xf0f17a/tau/internal/value"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTick_New(t *testing.T) {
	inst := "InstrumentName"
	price := value.NewFloat(343)
	timeStamp := time.Now()
	var count uint64 = 1
	tck := tick.New(instr.New(inst), price, timeStamp, count)
	assert.Equal(t, inst, tck.Instr.String())
	assert.Equal(t, price, tck.Price)
	assert.Equal(t, timeStamp, tck.Time)
	assert.Equal(t, count, tck.Count)
}
