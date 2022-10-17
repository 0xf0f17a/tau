package instr

type Instrument string

func New(v string) Instrument {
	return Instrument(v)
}

func (i Instrument) String() string {
	return string(i)
}

type Reader interface {
	ReadInstr() (Instrument, error)
}
