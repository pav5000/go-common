package errors

import (
	"errors"
	"fmt"
	"io"
)

var (
	New = errors.New
	As  = errors.As
	Is  = errors.Is
)

type wrappedError struct {
	err  error
	text string
}

func Wrp(err error, text string) error {
	if err == nil {
		return nil
	}

	return &wrappedError{
		err:  err,
		text: text,
	}
}

func (w *wrappedError) Error() string {
	if w == nil {
		return ""
	}

	return w.text + ": " + w.err.Error()
}

func (w *wrappedError) Unwrap() error {
	if w == nil {
		return nil
	}

	return w.err
}

func (w *wrappedError) Format(s fmt.State, _ rune) {
	_, _ = io.WriteString(s, w.text)
	_, _ = io.WriteString(s, ": ")
	_, _ = io.WriteString(s, w.err.Error())
}
