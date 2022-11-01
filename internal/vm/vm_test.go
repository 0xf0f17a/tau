package vm_test

import (
	"github.com/0xf0f17a/tau/internal/session"
	"github.com/0xf0f17a/tau/internal/value"
	"github.com/0xf0f17a/tau/internal/vm"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	open       = value.NewFloat(10)
	greenClose = value.NewFloat(15)
	high       = value.NewFloat(20)
	low        = value.NewFloat(5)
	ses        = session.New(open, greenClose, high, low)
)

func TestVMExecute(t *testing.T) {
	testCases := map[string]struct {
		opcodes []vm.Opcode
		signal  vm.Signal
	}{
		"singleHalt": {
			opcodes: []vm.Opcode{vm.OpCodeHalt},
			signal:  vm.SignalHold,
		},
	}
	for tn, tC := range testCases {
		t.Run(tn, func(t *testing.T) {
			v := vm.NewVM()
			for _, opcode := range tC.opcodes {
				v.Emit(opcode)
			}
			assert.Equal(t, tC.signal, v.Step(ses))
		})
	}
}

func TestVMHaltStopsVM(t *testing.T) {
	v := vm.NewVM()
	v.Emit(vm.OpCodeHalt)
	v.Step(ses)
	assert.False(t, v.IsRunning())
}

func TestVMOpenLoadsOpen(t *testing.T) {
	v := vm.NewVM()
	v.Emit(vm.OpCodeOpen)
	v.Step(ses)
	assert.Equal(t, open, v.RegAt(vm.RegA))
}

func TestVMCloseLoadsClose(t *testing.T) {
	v := vm.NewVM()
	v.Emit(vm.OpCodeClose)
	v.Step(ses)
	assert.Equal(t, greenClose, v.RegAt(vm.RegA))
}

func TestVMHighLoadsHigh(t *testing.T) {
	v := vm.NewVM()
	v.Emit(vm.OpCodeHigh)
	v.Step(ses)
	assert.Equal(t, high, v.RegAt(vm.RegA))
}

func TestVMLowLoadsLow(t *testing.T) {
	v := vm.NewVM()
	v.Emit(vm.OpCodeLow)
	v.Step(ses)
	assert.Equal(t, low, v.RegAt(vm.RegA))
}
