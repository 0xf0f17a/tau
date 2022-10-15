package tick

import (
	"github.com/0xf0f17a/tau/internal/instr"
	"github.com/0xf0f17a/tau/internal/value"
	"time"
)

type Tick struct {
	Instr instr.Instrument
	Price value.Value
	Time  time.Time
	Count uint64
}

func New(
	instr instr.Instrument,
	price value.Value,
	time time.Time,
	count uint64,
) *Tick {
	return &Tick{instr, price, time, count}
}
