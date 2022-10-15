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
})