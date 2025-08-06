package slice

import (
	"reflect"
	"strconv"

	"github.com/pav5000/go-common/types"
)

type (
	AnySlice[T any]               []T
	ComparableSlice[T comparable] []T
)

// New creates the new comparable slice from provided elements.
func New[T comparable](elems ...T) ComparableSlice[T] {
	return elems
}

// New creates the new "any" slice from provided elements.
func NewAny[T any](elems ...T) AnySlice[T] {
	return elems
}

// Copy creates another slice from the same elements.
func (s AnySlice[T]) Copy() AnySlice[T] {
	newSlice := make(AnySlice[T], len(s))
	copy(newSlice, s)

	return newSlice
}

// Copy creates another slice from the same elements.
func (s ComparableSlice[T]) Copy() ComparableSlice[T] {
	newSlice := make(ComparableSlice[T], len(s))
	copy(newSlice, s)

	return newSlice
}

// UniqueInplace leaves only unique elements
// you must use the returned value to get the result
// may mutate the array the slice is using.
func (s ComparableSlice[T]) UniqueInplace() ComparableSlice[T] {
	if len(s) < 2 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s[:1]
		}

		return s
	}

	m := make(map[T]struct{})

	newSlice := s[:0]
	for _, elem := range s {
		if _, ok := m[elem]; ok {
			continue
		}
		m[elem] = struct{}{}
		newSlice = append(newSlice, elem)
	}

	return newSlice
}

// RemoveElementsInplace removes all elements which value matches the provided value
// will mutate the array the slice is using.
func (s ComparableSlice[T]) RemoveElementsInplace(elemToRemove T) ComparableSlice[T] {
	newSlice := s[:0]
	for _, sliceElem := range s {
		if sliceElem == elemToRemove {
			continue
		}
		newSlice = append(newSlice, sliceElem)
	}

	return newSlice
}

// Unique leaves only unique elements
// returns the independent copy of the slice's data
// if the capacity is set to -1, len(s) will be used as the capacity.
func (s ComparableSlice[T]) Unique(capacity int) ComparableSlice[T] {
	switch len(s) {
	case 0:
		return nil
	case 1:
		return []T{s[0]}
	case 2:
		if s[0] == s[1] {
			return []T{s[0]}
		}

		return []T{s[0], s[1]}
	}

	if capacity < 0 {
		capacity = len(s)
	}
	m := make(map[T]struct{}, capacity)
	newSlice := make([]T, 0, capacity)

	for _, elem := range s {
		if _, ok := m[elem]; ok {
			continue
		}
		m[elem] = struct{}{}
		newSlice = append(newSlice, elem)
	}

	return newSlice
}

// IntSliceToStrings converts the slice of integers to the slice of strings.
func IntSliceToStrings[T types.Integer](slice []T) ComparableSlice[string] {
	if len(slice) == 0 {
		return nil
	}

	strs := make([]string, len(slice))
	switch reflect.TypeOf(slice[0]).Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		for i, value := range slice {
			strs[i] = strconv.FormatInt(int64(value), 10)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		for i, value := range slice {
			strs[i] = strconv.FormatUint(uint64(value), 10)
		}
	}

	return strs
}

// StringSliceToInt converts the slice of strings to the slice of integers.
func StringSliceToInt[T types.Integer](slice []string, defaultValue T) ComparableSlice[T] {
	if len(slice) == 0 {
		return nil
	}

	ints := make([]T, len(slice))
	t := reflect.TypeOf(ints[0])
	bitSize := t.Bits()
	switch t.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		for i, str := range slice {
			num, err := strconv.ParseInt(str, 10, bitSize)
			if err != nil {
				ints[i] = defaultValue

				continue
			}
			ints[i] = T(num)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		for i, str := range slice {
			num, err := strconv.ParseUint(str, 10, bitSize)
			if err != nil {
				ints[i] = defaultValue

				continue
			}
			ints[i] = T(num)
		}
	}

	return ints
}

// Has checks if the elem exists in the slice.
func (s ComparableSlice[T]) Has(elem T) bool {
	for _, sliceElem := range s {
		if elem == sliceElem {
			return true
		}
	}

	return false
}
