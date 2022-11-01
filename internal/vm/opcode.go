package vm

type Opcode int

const (
	OpCodeHalt  Opcode = 1
	OpCodeOpen  Opcode = 2
	OpCodeClose Opcode = 3
	OpCodeHigh  Opcode = 4
	OpCodeLow   Opcode = 5
)
