package value

type Value interface {
	Name() string
	Plus(another Value) Value
}

func NewFloat(v float32) Value {
	return Float(v)
}
