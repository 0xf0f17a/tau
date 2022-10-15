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
