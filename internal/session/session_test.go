package session_test

import (
	"github.com/0xf0f17a/tau/internal/session"
	"github.com/0xf0f17a/tau/internal/value"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session", func() {
	open := value.NewFloat(32)
	greenClose := value.NewFloat(43)
	redClose := value.NewFloat(26)
	high := value.NewFloat(51)
	low := value.NewFloat(23)
	greenSession := session.New(open, greenClose, high, low)
	redSession := session.New(open, redClose, high, low)
	It("session has open", func() {
		Expect(greenSession.Open).To(Equal(open))
		Expect(redSession.Open).To(Equal(open))
	})
	It("session has close", func() {
		Expect(greenSession.Close).To(Equal(greenClose))
		Expect(redSession.Close).To(Equal(redClose))
	})
	It("session has high", func() {
		Expect(greenSession.High).To(Equal(high))
		Expect(redSession.High).To(Equal(high))
	})
	It("session has low", func() {
		Expect(greenSession.Low).To(Equal(low))
		Expect(redSession.Low).To(Equal(low))
	})
	It("session can be bullish", func() {
		Expect(greenSession.IsBullish()).To(BeTrue())
		Expect(redSession.IsBullish()).To(BeFalse())
	})
	It("session can be bearish", func() {
		Expect(greenSession.IsBearish()).To(BeFalse())
		Expect(redSession.IsBearish()).To(BeTrue())
	})
	It("session has length", func() {
		Expect(greenSession.Length()).To(Equal(value.NewFloat(28)))
	})
	It("session has real body", func() {
		Expect(greenSession.RealBody()).To(Equal(value.NewFloat(11)))
		Expect(redSession.RealBody()).To(Equal(value.NewFloat(6)))
	})
	It("session has upper shadow", func() {
		Expect(greenSession.UpperShadow()).To(Equal(value.NewFloat(8)))
		Expect(redSession.UpperShadow()).To(Equal(value.NewFloat(19)))
	})
	It("session has lower shadow", func() {
		Expect(greenSession.LowerShadow()).To(Equal(value.NewFloat(9))) // 20
		Expect(redSession.LowerShadow()).To(Equal(value.NewFloat(3)))   // 9
	})
})
