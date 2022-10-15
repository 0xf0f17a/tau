package order

type Type int

func (t Type) Alternate() Type {
	switch t {
	case Buy:
		return Sell
	case Sell:
		return Buy
	default:
		return Sell
	}
}

const (
	Buy  Type = 1
	Sell Type = 2
)
