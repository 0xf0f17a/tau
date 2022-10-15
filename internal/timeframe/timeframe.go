package timeframe

type TimeFrame uint64

const (
	Minute uint64 = 1
	Hour          = 60 * Minute
)

func PerMinute() TimeFrame {
	return newDuration(1, Minute)
}

func PerNMinute(n uint8) TimeFrame {
	return newDuration(uint64(n), Minute)
}

func PerHour() TimeFrame {
	return newDuration(1, Hour)
}

func PerNHours(n uint8) TimeFrame {
	return newDuration(uint64(n), Hour)
}

func newDuration(multiple uint64, period uint64) TimeFrame {
	return TimeFrame(multiple * period)
}

func (tf TimeFrame) Duration() uint64 {
	return uint64(tf)
}
