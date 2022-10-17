package session_test

import (
	"github.com/0xf0f17a/tau/internal/session"
	"github.com/0xf0f17a/tau/internal/value"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	open         = value.NewFloat(32)
	greenClose   = value.NewFloat(43)
	redClose     = value.NewFloat(26)
	high         = value.NewFloat(51)
	low          = value.NewFloat(23)
	greenSession = session.New(open, greenClose, high, low)
	redSession   = session.New(open, redClose, high, low)
)

func TestSessionHasOpen(t *testing.T) {
	assert.Equal(t, open, greenSession.Open)
	assert.Equal(t, open, redSession.Open)
}

func TestSessionHasClose(t *testing.T) {
	assert.Equal(t, greenClose, greenSession.Close)
	assert.Equal(t, redClose, redSession.Close)
}

func TestSessionHasHigh(t *testing.T) {
	assert.Equal(t, high, greenSession.High)
	assert.Equal(t, high, redSession.High)
}

func TestSessionHasLow(t *testing.T) {
	assert.Equal(t, low, greenSession.Low)
	assert.Equal(t, low, redSession.Low)
}

func TestSession_IsBullish(t *testing.T) {
	assert.True(t, greenSession.IsBullish())
	assert.False(t, redSession.IsBullish())
}

func TestSession_IsBearish(t *testing.T) {
	assert.False(t, greenSession.IsBearish())
	assert.True(t, redSession.IsBearish())
}

func TestSession_Length(t *testing.T) {
	length := value.NewFloat(28)
	assert.Equal(t, length, greenSession.Length())
	assert.Equal(t, length, redSession.Length())
}

func TestSession_RealBody(t *testing.T) {
	assert.Equal(t, value.NewFloat(11), greenSession.RealBody())
	assert.Equal(t, value.NewFloat(6), redSession.RealBody())
}

func TestSession_LowerShadow(t *testing.T) {
	assert.Equal(t, value.NewFloat(9), greenSession.LowerShadow())
	assert.Equal(t, value.NewFloat(3), redSession.LowerShadow())
}

func TestSession_UpperShadow(t *testing.T) {
	assert.Equal(t, value.NewFloat(8), greenSession.UpperShadow())
	assert.Equal(t, value.NewFloat(19), redSession.UpperShadow())
}
