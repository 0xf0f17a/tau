package vm_test

import (
	"github.com/0xf0f17a/tau/internal/session"
	"github.com/0xf0f17a/tau/internal/value"
	"github.com/0xf0f17a/tau/internal/vm"
	"github.com/stretchr/testify/assert"
	"testing"
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
			open := value.NewFloat(10)
			greenClose := value.NewFloat(15)
			high := value.NewFloat(20)
			low := value.NewFloat(5)
			ses := session.New(open, greenClose, high, low)
			assert.Equal(t, tC.signal, v.Step(ses))
		})
	}
}

func TestVMHaltStopsVM(t *testing.T) {
	v := vm.NewVM()
	v.Emit(vm.OpCodeHalt)
	open := value.NewFloat(10)
	greenClose := value.NewFloat(15)
	high := value.NewFloat(20)
	low := value.NewFloat(5)
	ses := session.New(open, greenClose, high, low)
	v.Step(ses)
	assert.False(t, v.IsRunning())
}
