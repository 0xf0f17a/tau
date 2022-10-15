package timeframe

import "time"

type TimeFrame time.Duration

func PerMinute() TimeFrame {
	return TimeFrame(1 * time.Minute)
}

func (tf TimeFrame) Duration() time.Duration {
	return time.Duration(tf)
}
