package value

type Float float32

func (v Float) Name() string {
	return "Float"
}

func (v Float) Plus(another Value) Value {
	var float = another.(Float)
	return NewFloat(float32(v) + float32(float))
}
