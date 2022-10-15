package instr_test

import (
	"github.com/0xf0f17a/tau/internal/instr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Instr", func() {
	i := instr.New("InstrumentName")
	It("Instrument has ToString method", func() {
		Expect(i.String()).To(Equal("InstrumentName"))
	})
})
