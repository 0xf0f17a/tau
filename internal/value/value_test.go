package value_test

import (
	"github.com/0xf0f17a/tau/internal/value"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueHasName(t *testing.T) {
	testCases := map[string]struct {
		val      value.Value
		expected string
	}{
		"Float": {
			val:      value.NewFloat(23),
			expected: "Float",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.val.Name())
		})
	}
}

func TestValueCanBeAdded(t *testing.T) {
	testCases := map[string]struct {
		first  value.Value
		second value.Value
		result value.Value
	}{
		"FloatValid": {
			first:  value.NewFloat(23),
			second: value.NewFloat(45),
			result: value.NewFloat(68),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, tc.first.Plus(tc.second))
		})
	}
}

func TestValueCanBeSubtracted(t *testing.T) {
	testCases := map[string]struct {
		first  value.Value
		second value.Value
		result value.Value
	}{
		"FloatValid": {
			first:  value.NewFloat(45),
			second: value.NewFloat(23),
			result: value.NewFloat(22),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, tc.first.Minus(tc.second))
		})
	}
}

func TestValueCanBeMultiplied(t *testing.T) {
	testCases := map[string]struct {
		first  value.Value
		second value.Value
		result value.Value
	}{
		"FloatValid": {
			first:  value.NewFloat(23),
			second: value.NewFloat(45),
			result: value.NewFloat(1035),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, tc.first.Multiply(tc.second))
		})
	}
}

func TestValueCanBeDivided(t *testing.T) {
	testCases := map[string]struct {
		first  value.Value
		second value.Value
		result value.Value
	}{
		"FloatValid": {
			first:  value.NewFloat(46),
			second: value.NewFloat(23),
			result: value.NewFloat(2),
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, tc.first.Divide(tc.second))
		})
	}
}

func TestValueHasGreaterThanComparison(t *testing.T) {
	testCases := map[string]struct {
		first        value.Value
		second       value.Value
		third        value.Value
		first2Second bool
		first2Third  bool
		second2Third bool
	}{
		"FloatValid": {
			first:        value.NewFloat(46),
			second:       value.NewFloat(23),
			third:        value.NewFloat(23),
			first2Second: true,
			first2Third:  true,
			second2Third: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.first2Second, tc.first.IsGreaterThan(tc.second))
			assert.Equal(t, tc.first2Third, tc.first.IsGreaterThan(tc.third))
			assert.Equal(t, tc.second2Third, tc.second.IsGreaterThan(tc.third))
		})
	}
}

func TestValueHasLessThanComparison(t *testing.T) {
	testCases := map[string]struct {
		first        value.Value
		second       value.Value
		third        value.Value
		first2Second bool
		first2Third  bool
		second2Third bool
	}{
		"FloatValid": {
			first:        value.NewFloat(46),
			second:       value.NewFloat(23),
			third:        value.NewFloat(23),
			first2Second: false,
			first2Third:  false,
			second2Third: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.first2Second, tc.first.IsLessThan(tc.second))
			assert.Equal(t, tc.first2Third, tc.first.IsLessThan(tc.third))
			assert.Equal(t, tc.second2Third, tc.second.IsLessThan(tc.third))
		})
	}
}

func TestValueHasGreaterThanEqualComparison(t *testing.T) {
	testCases := map[string]struct {
		first        value.Value
		second       value.Value
		third        value.Value
		first2Second bool
		first2Third  bool
		second2Third bool
	}{
		"FloatValid": {
			first:        value.NewFloat(46),
			second:       value.NewFloat(23),
			third:        value.NewFloat(23),
			first2Second: true,
			first2Third:  true,
			second2Third: true,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.first2Second, tc.first.IsGreaterThanEqualTo(tc.second))
			assert.Equal(t, tc.first2Third, tc.first.IsGreaterThanEqualTo(tc.third))
			assert.Equal(t, tc.second2Third, tc.second.IsGreaterThanEqualTo(tc.third))
		})
	}
}

func TestValueHasLessThanEqualComparison(t *testing.T) {
	testCases := map[string]struct {
		first        value.Value
		second       value.Value
		third        value.Value
		first2Second bool
		first2Third  bool
		second2Third bool
	}{
		"FloatValid": {
			first:        value.NewFloat(46),
			second:       value.NewFloat(23),
			third:        value.NewFloat(23),
			first2Second: false,
			first2Third:  false,
			second2Third: true,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.first2Second, tc.first.IsLessThanEqualTo(tc.second))
			assert.Equal(t, tc.first2Third, tc.first.IsLessThanEqualTo(tc.third))
			assert.Equal(t, tc.second2Third, tc.second.IsLessThanEqualTo(tc.third))
		})
	}
}

func TestValueHasEqualComparison(t *testing.T) {
	testCases := map[string]struct {
		first        value.Value
		second       value.Value
		third        value.Value
		first2Second bool
		first2Third  bool
		second2Third bool
	}{
		"FloatValid": {
			first:        value.NewFloat(46),
			second:       value.NewFloat(23),
			third:        value.NewFloat(23),
			first2Second: false,
			first2Third:  false,
			second2Third: true,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.first2Second, tc.first.IsEqualTo(tc.second))
			assert.Equal(t, tc.first2Third, tc.first.IsEqualTo(tc.third))
			assert.Equal(t, tc.second2Third, tc.second.IsEqualTo(tc.third))
		})
	}
}

func TestValueHasNotEqualComparison(t *testing.T) {
	testCases := map[string]struct {
		first        value.Value
		second       value.Value
		third        value.Value
		first2Second bool
		first2Third  bool
		second2Third bool
	}{
		"FloatValid": {
			first:        value.NewFloat(46),
			second:       value.NewFloat(23),
			third:        value.NewFloat(23),
			first2Second: true,
			first2Third:  true,
			second2Third: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.first2Second, tc.first.IsNotEqualTo(tc.second))
			assert.Equal(t, tc.first2Third, tc.first.IsNotEqualTo(tc.third))
			assert.Equal(t, tc.second2Third, tc.second.IsNotEqualTo(tc.third))
		})
	}
}
