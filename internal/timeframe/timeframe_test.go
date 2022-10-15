package timeframe_test

import (
	"github.com/0xf0f17a/tau/internal/timeframe"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Timeframe", func() {
	It("Has Per minute timeframe", func() {
		tf := timeframe.PerMinute()
		Expect(tf.Duration()).To(Equal(1 * timeframe.Minute))
	})
	It("Has Per N minute timeframe", func() {
		tf := timeframe.PerNMinute(15)
		Expect(tf.Duration()).To(Equal(15 * timeframe.Minute))
	})
	It("Has Per Hour timeframe", func() {
		tf := timeframe.PerHour()
		Expect(tf.Duration()).To(Equal(1 * timeframe.Hour))
	})
	It("Has Per N Hour timeframe", func() {
		tf := timeframe.PerNHours(15)
		Expect(tf.Duration()).To(Equal(15 * timeframe.Hour))
	})
	It("Has Per Day timeframe", func() {
		tf := timeframe.PerDay()
		Expect(tf.Duration()).To(Equal(1 * timeframe.Day))
	})
	It("Has Per N Day timeframe", func() {
		tf := timeframe.PerNDays(15)
		Expect(tf.Duration()).To(Equal(15 * timeframe.Day))
	})
	It("Has Per Week timeframe", func() {
		tf := timeframe.PerWeek()
		Expect(tf.Duration()).To(Equal(1 * timeframe.Week))
	})
	It("Has Per N Week timeframe", func() {
		tf := timeframe.PerNWeeks(15)
		Expect(tf.Duration()).To(Equal(15 * timeframe.Week))
	})
	It("Has Per Month timeframe", func() {
		tf := timeframe.PerMonth()
		Expect(tf.Duration()).To(Equal(1 * timeframe.Month))
	})
	It("Has Per N Month timeframe", func() {
		tf := timeframe.PerNMonths(15)
		Expect(tf.Duration()).To(Equal(15 * timeframe.Month))
	})
})
