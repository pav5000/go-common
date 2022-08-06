package slice

import (
	"strconv"

	"github.com/pav5000/go-common/types"
)

type AnySlice[T any] []T
type ComparableSlice[T comparable] []T

// New creates the new comparable slice from provided elements
func New[T comparable](elems ...T) ComparableSlice[T] {
	return elems
}

// New creates the new "any" slice from provided elements
func NewAny[T any](elems ...T) AnySlice[T] {
	return elems
}

// Copy creates another slice from the same elements
func (s AnySlice[T]) Copy() AnySlice[T] {
	newSlice := make(AnySlice[T], len(s))
	for i, elem := range s {
		newSlice[i] = elem
	}
	return newSlice
}

// Copy creates another slice from the same elements
func (s ComparableSlice[T]) Copy() ComparableSlice[T] {
	newSlice := make(ComparableSlice[T], len(s))
	for i, elem := range s {
		newSlice[i] = elem
	}
	return newSlice
}

// UniqueInplace leaves only unique elements
// you must use the returned value to get the result
// may mutate the array the slice is using
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

// Unique leaves only unique elements
// returns the independent copy of the slice's data
// if the capacity is set to -1, len(s) will be used as the capacity
func (s ComparableSlice[T]) Unique(cap int) ComparableSlice[T] {
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

	if cap < 0 {
		cap = len(s)
	}
	m := make(map[T]struct{}, cap)
	newSlice := make([]T, 0, cap)

	for _, elem := range s {
		if _, ok := m[elem]; ok {
			continue
		}
		m[elem] = struct{}{}
		newSlice = append(newSlice, elem)
	}
	return newSlice
}

// IntSliceToStrings converts the slice of signed integers to the slice of strings
func IntSliceToStrings[T types.SignedInteger](slice []T) []string {
	if len(slice) == 0 {
		return nil
	}

	strs := make([]string, len(slice))
	for i, value := range slice {
		strs[i] = strconv.FormatInt(int64(value), 10)
	}
	return strs
}

// UintSliceToStrings converts the slice of unsigned integers to the slice of strings
func UintSliceToStrings[T types.UnsignedInteger](slice []T) []string {
	if len(slice) == 0 {
		return nil
	}

	strs := make([]string, len(slice))
	for i, value := range slice {
		strs[i] = strconv.FormatUint(uint64(value), 10)
	}
	return strs
}
