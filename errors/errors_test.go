package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WrappedError_IncludesBothTextFromErrorAndTextFromTextField(t *testing.T) {
	originalError := errors.New("original error")

	wrappedError := Wrp(originalError, "some text")

	assert.ErrorContains(t, wrappedError, "original error")
	assert.ErrorContains(t, wrappedError, "some text")
}

func Test_WrappedError_SupportsUnwrapChain(t *testing.T) {
	originalError := errors.New("original error")

	wrappedError := Wrp(originalError, "some text")

	assert.True(t, errors.Is(wrappedError, originalError))
}

func Test_WrappedError_WorksWithStreamingFormatter(t *testing.T) {
	originalError := errors.New("original error")

	wrappedError := Wrp(originalError, "some text")

	assert.Equal(t, "some text: original error", fmt.Sprintf("%s", wrappedError))
}
