package vm

import (
	"github.com/0xf0f17a/tau/internal/session"
	"github.com/0xf0f17a/tau/internal/value"
)

type VM struct {
	codes     []Opcode
	registers [Max]value.Value
	running   bool
}

func NewVM() *VM {
	return &VM{}
}

func (v *VM) Emit(op Opcode) {
	v.codes = append(v.codes, op)
}

func (v *VM) Step(sess *session.Session) Signal {
	v.running = true
	for _, code := range v.codes {
		switch code {
		case OpCodeHalt:
			v.running = false
		case OpCodeOpen:
			v.registers[RegA] = sess.Open
		default:
			break
		}
	}
	return SignalHold
}

func (v *VM) IsRunning() bool {
	return v.running
}

func (v *VM) RegAt(i Reg) value.Value {
	return v.registers[i]
}
