package errors_test

import (
	"fmt"
	"github.com/0xf0f17a/tau/internal/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoArgs(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			assert.Fail(t, "E() did not panic")
		}
	}()
	_ = errors.E()
}

func TestErrorNewString(t *testing.T) {
	errString := "unknown error occurred"
	err := errors.Str(errString)
	assert.Equal(t, errString, err.Error())
}

func TestErrorFmt(t *testing.T) {
	errString := fmt.Sprintf("Hello this is %v", 43)
	err := errors.Fmt("Hello this is %v", 43)
	assert.Equal(t, errString, err.Error())
}

func TestErrorString(t *testing.T) {
	const op = errors.Op("account.New")
	underlying := errors.Str("unknown error")
	assert.Equal(t, "account.New: unknown error", errors.E(op, underlying).Error())
}

func TestErrorAsErr(t *testing.T) {
	const op = errors.Op("account.New")
	underlying := errors.Str("unknown error")
	err := errors.E(op, underlying)
	actual := errors.AsErr(err)
	assert.Equal(t, op, actual.Op)
	assert.Equal(t, underlying, actual.Err)
}
