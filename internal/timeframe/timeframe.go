package timeframe

import "time"

type TimeFrame time.Duration

func PerMinute() TimeFrame {
	return PerNMinute(1)
}

func PerNMinute(n uint8) TimeFrame {
	return TimeFrame(time.Duration(n) * time.Minute)
}

func (tf TimeFrame) Duration() time.Duration {
	return time.Duration(tf)
}
