package timeframe_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTimeframe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Timeframe Suite")
}
