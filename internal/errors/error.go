package errors

import (
	"bytes"
	"errors"
	"fmt"
)

type Op string

func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errors.E with no arguments")
	}
	err := &Err{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			err.Op = arg
		case *Err:
			// Make a copy
			c := *arg
			err.Err = &c
		case error:
			err.Err = arg
		default:
			//
		}
	}
	return err
}
func AsErr(err error) *Err {
	var e *Err
	errors.As(err, &e)
	return e
}

func Str(msg string) error {
	return &errorString{msg: msg}
}

func Fmt(format string, args ...interface{}) error {
	return &errorString{msg: fmt.Sprintf(format, args...)}
}

type Err struct {
	Op  Op
	Err error
}

func (e *Err) Error() string {
	b := new(bytes.Buffer)
	if e.Op != "" {
		b.WriteString(string(e.Op))
	}
	if e.Err != nil {
		pad(b, ": ")
		b.WriteString(e.Err.Error())
	}
	if b.Len() == 0 {
		return "no error"
	}
	return b.String()
}

type errorString struct {
	msg string
}

func (es *errorString) Error() string {
	return es.msg
}

func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}
	b.WriteString(str)
}
