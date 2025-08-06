package slice

import "math/rand"

// RandomElement selects a random element from the slice
// returns default value if slice is empty
// don't use in cryptographic applications (using math/rand here).
func RandomElement[T any](slice []T) T {
	if len(slice) == 0 {
		var elem T

		return elem
	}

	return slice[rand.Intn(len(slice))]
}
