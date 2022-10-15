package session

import "github.com/0xf0f17a/tau/internal/value"

type Session struct {
	Open  value.Value
	Close value.Value
	High  value.Value
	Low   value.Value
}

func New(open value.Value, close value.Value, high value.Value, low value.Value) *Session {
	return &Session{open, close, high, low}
}

func (s *Session) IsBullish() bool {
	return s.Open.IsLessThan(s.Close)
}

func (s *Session) IsBearish() bool {
	return s.Open.IsGreaterThan(s.Close)
}

func (s *Session) Length() value.Value {
	return s.High.Minus(s.Low)
}

func (s *Session) RealBody() value.Value {
	if s.IsBullish() {
		return s.Close.Minus(s.Open)
	} else {
		return s.Open.Minus(s.Close)
	}
}

func (s *Session) UpperShadow() value.Value {
	if s.IsBullish() {
		return s.High.Minus(s.Close)
	} else {
		return s.High.Minus(s.Open)
	}
}

func (s *Session) LowerShadow() value.Value {
	if s.IsBullish() {
		return s.Open.Minus(s.Low)
	} else {
		return s.Close.Minus(s.Low)
	}
}
