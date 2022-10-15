package timeframe

import "time"

type TimeFrame time.Duration

func PerMinute() TimeFrame {
	return newDuration(1, time.Minute)
}

func PerNMinute(n uint8) TimeFrame {
	return newDuration(n, time.Minute)
}

func PerHour() TimeFrame {
	return newDuration(1, time.Hour)
}

func PerNHours(n uint8) TimeFrame {
	return newDuration(n, time.Hour)
}

func newDuration(multiple uint8, period time.Duration) TimeFrame {
	return TimeFrame(time.Duration(multiple) * period)
}

func (tf TimeFrame) Duration() time.Duration {
	return time.Duration(tf)
}
