package stringutil

// Head leaves no more than letterCount unicode letters
// from the beginning of the string.
// Trimming the end of the string if needed
func Head(str string, letterCount int) string {
	if len(str) <= letterCount {
		return str
	}
	count := 0
	for i := range str {
		if count == letterCount {
			return str[:i]
		}
		count++
	}
	return str
}

// HeadBytes leaves no more than byteCount bytes
// from the beginning of the slice.
// Trimming the end if needed
func HeadBytes(data []byte, byteCount int) []byte {
	if len(data) <= byteCount {
		return data
	}
	return data[:byteCount]
}

// TailBytes leaves no more than byteCount bytes
// from the end of the slice.
// Trimming the beginning if needed.
func TailBytes(data []byte, byteCount int) []byte {
	if len(data) <= byteCount {
		return data
	}
	return data[len(data)-byteCount:]
}
