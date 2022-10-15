package timeframe

type TimeFrame uint64

const (
	Minute uint64 = 1
	Hour          = 60 * Minute
	Day           = 24 * Hour
	Week          = 7 * Day
	Month         = 4 * Week
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

func PerDay() TimeFrame {
	return newDuration(1, Day)
}

func PerNDays(n uint8) TimeFrame {
	return newDuration(uint64(n), Day)
}

func PerWeek() TimeFrame {
	return newDuration(1, Week)
}

func PerNWeeks(n uint8) TimeFrame {
	return newDuration(uint64(n), Week)
}

func PerMonth() TimeFrame {
	return newDuration(1, Month)
}

func PerNMonths(n uint8) TimeFrame {
	return newDuration(uint64(n), Month)
}

func newDuration(multiple uint64, period uint64) TimeFrame {
	return TimeFrame(multiple * period)
}

func (tf TimeFrame) Duration() uint64 {
	return uint64(tf)
}
