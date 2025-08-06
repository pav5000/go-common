package stringutil

import (
	"math/rand"
	"strings"
)

const (
	LowercaseLatinLetters   = "abcdefghijklmnopqrstuvwxyz"
	UppercaseLatinLetters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LatinCharacters         = LowercaseLatinLetters + UppercaseLatinLetters
	Digits                  = "0123456789"
	LatinCaractersAndDigits = LatinCharacters + Digits
)

// RandomASCIIString generates a random string from characters which present in the provided charset
// only 1-byte characters are supported.
func RandomASCIIString(charset string, length int) string {
	res := strings.Builder{}
	res.Grow(length)
	for range length {
		res.WriteByte(charset[rand.Intn(len(charset))])
	}

	return res.String()
}
