package instr_test

import (
	"github.com/0xf0f17a/tau/internal/errors"
	"github.com/0xf0f17a/tau/internal/instr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstrument_String(t *testing.T) {
	inst := "InstrumentName"
	i := instr.New(inst)
	assert.Equal(t, inst, i.String())
}

type ConstantInstrumentReader struct {
	str string
	err error
}

func (r *ConstantInstrumentReader) ReadInstr() (instr.Instrument, error) {
	if r.err != nil {
		return "", r.err
	} else {
		return instr.Instrument(r.str), nil
	}
}

func (r *ConstantInstrumentReader) WriteInstr(instr.Instrument) error {
	return r.err
}

func TestInstrumentReader_ReadInstr_Success(t *testing.T) {
	const str = "InstrumentName"
	testCases := map[string]instr.Reader{
		"constantInstrumentReader": &ConstantInstrumentReader{str: str, err: nil},
	}
	for name, reader := range testCases {
		t.Run(name, func(t *testing.T) {
			inst, err := reader.ReadInstr()
			assert.NoError(t, err)
			assert.Equal(t, str, inst.String())
		})
	}
}

func TestInstrumentReader_ReadInstr_Failed(t *testing.T) {
	testCases := map[string]instr.Reader{
		"constantInstrumentReader": &ConstantInstrumentReader{
			str: "InstrumentName",
			err: errors.Str("unknown error"),
		},
	}
	for name, reader := range testCases {
		t.Run(name, func(t *testing.T) {
			inst, err := reader.ReadInstr()
			assert.Error(t, err)
			assert.Equal(t, instr.Instrument(""), inst)
		})
	}
}

func TestInstrumentReader_WriteInstr_Success(t *testing.T) {
	const str = "InstrumentName"
	testCases := map[string]instr.Writer{
		"constantInstrumentReader": &ConstantInstrumentReader{err: nil},
	}
	for name, writer := range testCases {
		t.Run(name, func(t *testing.T) {
			err := writer.WriteInstr(instr.New(str))
			assert.NoError(t, err)
		})
	}
}

func TestInstrumentReader_WriteInstr_Failed(t *testing.T) {
	const str = "InstrumentName"
	testCases := map[string]instr.Writer{
		"constantInstrumentReader": &ConstantInstrumentReader{err: errors.Str("unknown error")},
	}
	for name, writer := range testCases {
		t.Run(name, func(t *testing.T) {
			err := writer.WriteInstr(instr.New(str))
			assert.Error(t, err)
			assert.Equal(t, "unknown error", err.Error())
		})
	}
}
