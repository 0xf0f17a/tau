package value

type Float float32

func (v Float) Name() string {
	return "Float"
}

func (v Float) Plus(another Value) Value {
	return v + another.(Float)
}

func (v Float) Minus(another Value) Value {
	return v - another.(Float)
}

func (v Float) Multiply(another Value) Value {
	return v * another.(Float)
}

func (v Float) Divide(another Value) Value {
	return v / another.(Float)
}

func (v Float) IsGreaterThan(another Value) bool {
	return v > another.(Float)
}

func (v Float) IsLessThan(another Value) bool {
	return v < another.(Float)
}

func (v Float) IsGreaterThanEqualTo(another Value) bool {
	return v >= another.(Float)
}

func (v Float) IsLessThanEqualTo(another Value) bool {
	return v <= another.(Float)
}

func (v Float) IsEqualTo(another Value) bool {
	return v == another.(Float)
}

func (v Float) IsNotEqualTo(another Value) bool {
	return v != another.(Float)
}
