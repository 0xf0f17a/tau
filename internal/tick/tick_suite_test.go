package tick_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTick(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tick Suite")
}
