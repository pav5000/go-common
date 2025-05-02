package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Head(t *testing.T) {
	assert.Equal(t, "some te", Head("some text", 7))
	assert.Equal(t, "some tex", Head("some text", 8))
	assert.Equal(t, "some text", Head("some text", 20))
	assert.Equal(t, "some text", Head("some text", 9))

	assert.Equal(t, "", Head("привет", 0))
	assert.Equal(t, "п", Head("привет", 1))
	assert.Equal(t, "пр", Head("привет", 2))
	assert.Equal(t, "при", Head("привет", 3))
	assert.Equal(t, "прив", Head("привет", 4))
	assert.Equal(t, "приве", Head("привет", 5))
	assert.Equal(t, "привет", Head("привет", 6))
	assert.Equal(t, "привет", Head("привет", 7))
	assert.Equal(t, "привет", Head("привет", 20))

	assert.Equal(t, "", Head("", 10))
	assert.Equal(t, "", Head("", 0))
	assert.Equal(t, "a", Head("a", 10))
	assert.Equal(t, "", Head("a", 0))
}

func Test_HeadBytes(t *testing.T) {
	assert.Equal(t, []byte(""), HeadBytes([]byte("test"), 0))
	assert.Equal(t, []byte("t"), HeadBytes([]byte("test"), 1))
	assert.Equal(t, []byte("te"), HeadBytes([]byte("test"), 2))
	assert.Equal(t, []byte("tes"), HeadBytes([]byte("test"), 3))
	assert.Equal(t, []byte("test"), HeadBytes([]byte("test"), 4))
	assert.Equal(t, []byte("test"), HeadBytes([]byte("test"), 5))
	assert.Equal(t, []byte("test"), HeadBytes([]byte("test"), 50))
}

func Test_TailBytes(t *testing.T) {
	assert.Equal(t, []byte(""), TailBytes([]byte("test"), 0))
	assert.Equal(t, []byte("t"), TailBytes([]byte("test"), 1))
	assert.Equal(t, []byte("st"), TailBytes([]byte("test"), 2))
	assert.Equal(t, []byte("est"), TailBytes([]byte("test"), 3))
	assert.Equal(t, []byte("test"), TailBytes([]byte("test"), 4))
	assert.Equal(t, []byte("test"), TailBytes([]byte("test"), 5))
	assert.Equal(t, []byte("test"), TailBytes([]byte("test"), 50))
}
