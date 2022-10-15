package value

type Value interface {
	Name() string
	Plus(another Value) Value
	Minus(another Value) Value
	Multiply(another Value) Value
	Divide(another Value) Value
	IsGreaterThan(another Value) bool
	IsLessThan(another Value) bool
	IsGreaterThanEqualTo(another Value) bool
	IsLessThanEqualTo(another Value) bool
	IsEqualTo(another Value) bool
	IsNotEqualTo(another Value) bool
}

func NewFloat(v float32) Value {
	return Float(v)
}
