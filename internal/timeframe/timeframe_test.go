package timeframe_test

import (
	"github.com/0xf0f17a/tau/internal/timeframe"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Timeframe", func() {
	It("Has Per minute timeframe", func() {
		tf := timeframe.PerMinute()
		Expect(tf.Duration()).To(Equal(1 * time.Minute))
	})
})
