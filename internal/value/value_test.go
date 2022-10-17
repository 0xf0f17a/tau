package value_test

import (
	"github.com/0xf0f17a/tau/internal/value"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueHasName(t *testing.T) {
	assert.Equal(t, "Float", value.NewFloat(23).Name())
}

func TestValueCanBeAdded(t *testing.T) {
	first := value.NewFloat(23)
	second := value.NewFloat(45)
	result := value.NewFloat(68)
	assert.Equal(t, result, first.Plus(second))
}

func TestValueCanBeSubtracted(t *testing.T) {
	first := value.NewFloat(45)
	second := value.NewFloat(23)
	result := value.NewFloat(22)
	assert.Equal(t, result, first.Minus(second))
}

func TestValueCanBeMultiplied(t *testing.T) {
	first := value.NewFloat(23)
	second := value.NewFloat(45)
	result := value.NewFloat(1035)
	assert.Equal(t, result, first.Multiply(second))
}

func TestValueCanBeDivided(t *testing.T) {
	first := value.NewFloat(46)
	second := value.NewFloat(23)
	result := value.NewFloat(2)
	assert.Equal(t, result, first.Divide(second))
}

func TestValueHasGreaterThanComparison(t *testing.T) {
	first := value.NewFloat(46)
	second := value.NewFloat(23)
	third := value.NewFloat(23)
	assert.True(t, first.IsGreaterThan(second))
	assert.False(t, second.IsGreaterThan(third))
}

func TestValueHasLessThanComparison(t *testing.T) {
	first := value.NewFloat(46)
	second := value.NewFloat(23)
	third := value.NewFloat(23)
	assert.True(t, second.IsLessThan(first))
	assert.False(t, second.IsLessThan(third))
}
func TestValueHasGreaterThanEqualComparison(t *testing.T) {
	first := value.NewFloat(46)
	second := value.NewFloat(23)
	third := value.NewFloat(23)
	assert.True(t, first.IsGreaterThanEqualTo(second))
	assert.True(t, second.IsGreaterThanEqualTo(third))
}

func TestValueHasLessThanEqualComparison(t *testing.T) {
	first := value.NewFloat(46)
	second := value.NewFloat(23)
	third := value.NewFloat(23)
	assert.True(t, second.IsLessThanEqualTo(first))
	assert.True(t, second.IsLessThanEqualTo(third))
}

func TestValueHasEqualComparison(t *testing.T) {
	first := value.NewFloat(46)
	second := value.NewFloat(23)
	third := value.NewFloat(23)
	assert.False(t, first.IsEqualTo(second))
	assert.True(t, second.IsEqualTo(third))
}

func TestValueHasNotEqualComparison(t *testing.T) {
	first := value.NewFloat(46)
	second := value.NewFloat(23)
	third := value.NewFloat(23)
	assert.True(t, second.IsNotEqualTo(first))
	assert.False(t, second.IsNotEqualTo(third))
}
