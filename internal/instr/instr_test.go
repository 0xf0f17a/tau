package instr_test

import (
	"github.com/0xf0f17a/tau/internal/instr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstrument_String(t *testing.T) {
	inst := "InstrumentName"
	i := instr.New(inst)
	assert.Equal(t, inst, i.String())
}
