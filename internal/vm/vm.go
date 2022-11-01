package vm

import (
	"github.com/0xf0f17a/tau/internal/session"
)

type VM struct {
	codes   []Opcode
	running bool
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
		default:
			break
		}
	}
	return SignalHold
}

func (v *VM) IsRunning() bool {
	return v.running
}
