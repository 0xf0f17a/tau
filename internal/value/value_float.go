package value

type float float32

func (v float) Name() string {
	return "Float"
}

func (v float) Plus(another Value) Value {
	return v + another.(float)
}

func (v float) Minus(another Value) Value {
	return v - another.(float)
}

func (v float) Multiply(another Value) Value {
	return v * another.(float)
}

func (v float) Divide(another Value) Value {
	return v / another.(float)
}

func (v float) IsGreaterThan(another Value) bool {
	return v > another.(float)
}

func (v float) IsLessThan(another Value) bool {
	return v < another.(float)
}

func (v float) IsGreaterThanEqualTo(another Value) bool {
	return v >= another.(float)
}

func (v float) IsLessThanEqualTo(another Value) bool {
	return v <= another.(float)
}

func (v float) IsEqualTo(another Value) bool {
	return v == another.(float)
}

func (v float) IsNotEqualTo(another Value) bool {
	return v != another.(float)
}
