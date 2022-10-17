package timeframe_test

import (
	"github.com/0xf0f17a/tau/internal/timeframe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeFrame_Duration(t *testing.T) {
	testCases := map[string]struct {
		tf  timeframe.TimeFrame
		val uint64
	}{
		"PerMinute": {
			tf:  timeframe.PerMinute(),
			val: 1 * timeframe.Minute,
		},
		"PerNMinute": {
			tf:  timeframe.PerNMinute(15),
			val: 15 * timeframe.Minute,
		},
		"PerHour": {
			tf:  timeframe.PerHour(),
			val: 1 * timeframe.Hour,
		},
		"PerNHour": {
			tf:  timeframe.PerNHours(4),
			val: 4 * timeframe.Hour,
		},
		"PerDay": {
			tf:  timeframe.PerDay(),
			val: 1 * timeframe.Day,
		},
		"PerNDay": {
			tf:  timeframe.PerNDays(5),
			val: 5 * timeframe.Day,
		},
		"PerWeek": {
			tf:  timeframe.PerWeek(),
			val: 1 * timeframe.Week,
		},
		"PerNWeek": {
			tf:  timeframe.PerNWeeks(3),
			val: 3 * timeframe.Week,
		},
		"PerMonth": {
			tf:  timeframe.PerMonth(),
			val: 1 * timeframe.Month,
		},
		"PerNMonth": {
			tf:  timeframe.PerNMonths(2),
			val: 2 * timeframe.Month,
		},
		"PerQuarter": {
			tf:  timeframe.PerQuarter(),
			val: 1 * timeframe.Quarter,
		},
		"PerNQuarter": {
			tf:  timeframe.PerNQuarter(2),
			val: 2 * timeframe.Quarter,
		},
		"PerYear": {
			tf:  timeframe.PerYear(),
			val: 1 * timeframe.Year,
		},
		"PerNYear": {
			tf:  timeframe.PerNYear(4),
			val: 4 * timeframe.Year,
		},
	}
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			tf := tt.tf.Duration()
			tv := tt.val
			assert.Equal(t, tf, tv)
		})
	}
}
