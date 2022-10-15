package value

type Value interface {
	Name() string
}

func NewFloat(v float32) Value {
	return Float(v)
}
