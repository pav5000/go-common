package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Copy_Comparable(t *testing.T) {
	slice := New(1, 2, 3, 4)

	copied := slice.Copy()
	slice[0] = 10

	assert.EqualValues(t, []int{1, 2, 3, 4}, copied)
}

func Test_Copy_Any(t *testing.T) {
	slice := NewAny(1, 2, 3, 4)

	copied := slice.Copy()
	slice[0] = 10

	assert.EqualValues(t, []int{1, 2, 3, 4}, copied)
}

func Test_UniqueInplace_NoElements(t *testing.T) {
	slice := New[int]()

	slice = slice.UniqueInplace()

	assert.Len(t, slice, 0)
}

func Test_UniqueInplace_OneElement(t *testing.T) {
	slice := New(1)

	slice = slice.UniqueInplace()

	assert.EqualValues(t, []int{1}, slice)
}

func Test_UniqueInplace_TwoEqualElements(t *testing.T) {
	slice := New(1, 1)

	slice = slice.UniqueInplace()

	assert.EqualValues(t, []int{1}, slice)
}

func Test_UniqueInplace_TwoDifferentElements(t *testing.T) {
	slice := New(1, 2)

	slice = slice.UniqueInplace()

	assert.EqualValues(t, []int{1, 2}, slice)
}

func Test_UniqueInplace_SeveralElementsWithDuplicates(t *testing.T) {
	slice := New(1, 2, 1, 3, 4, 3, 2)

	slice = slice.UniqueInplace()

	assert.EqualValues(t, []int{1, 2, 3, 4}, slice)
}

func Test_UniqueInplace_SeveralUniqueElements(t *testing.T) {
	slice := New(1, 2, 3, 4)

	slice = slice.UniqueInplace()

	assert.EqualValues(t, []int{1, 2, 3, 4}, slice)
}

func Test_Unique_NoElements(t *testing.T) {
	slice := New[int]()

	slice = slice.Unique(0)

	assert.Len(t, slice, 0)
}

func Test_Unique_OneElement(t *testing.T) {
	slice := New(1)

	newSlice := slice.Unique(0)
	slice[0] = 10

	assert.EqualValues(t, []int{1}, newSlice)
}

func Test_Unique_TwoEqualElements(t *testing.T) {
	slice := New(1, 1)

	newSlice := slice.Unique(0)
	slice[0] = 10

	assert.EqualValues(t, []int{1}, newSlice)
}

func Test_Unique_TwoDifferentElements(t *testing.T) {
	slice := New(1, 2)

	newSlice := slice.Unique(0)
	slice[0] = 10

	assert.EqualValues(t, []int{1, 2}, newSlice)
}

func Test_Unique_SeveralElementsWithDuplicates(t *testing.T) {
	slice := New(1, 2, 1, 3, 4, 3, 2)

	newSlice := slice.Unique(4)
	slice[0] = 10

	assert.EqualValues(t, []int{1, 2, 3, 4}, newSlice)
}

func Test_Unique_SeveralUniqueElements(t *testing.T) {
	slice := New(1, 2, 3, 4)

	newSlice := slice.Unique(4)
	slice[0] = 10

	assert.EqualValues(t, []int{1, 2, 3, 4}, newSlice)
}

func Test_Unique_NegaviteCap(t *testing.T) {
	slice := New(1, 2, 3, 4)

	newSlice := slice.Unique(-1)
	slice[0] = 10

	assert.EqualValues(t, []int{1, 2, 3, 4}, newSlice)
}

func Test_IntSliceToStringsSigned(t *testing.T) {
	slice := New(-2, -1, 1, 2, 3, 4)

	newSlice := IntSliceToStrings(slice)

	assert.EqualValues(t, []string{"-2", "-1", "1", "2", "3", "4"}, newSlice)
}

func Test_UintSliceToStringsUnsigned(t *testing.T) {
	slice := []uint64{1, 2, 3, 4, 18446744073709551615}

	newSlice := IntSliceToStrings(slice)

	assert.EqualValues(t, []string{"1", "2", "3", "4", "18446744073709551615"}, newSlice)
}

func Test_StringSliceToInt_simple(t *testing.T) {
	slice := []string{"1", "2", "3", "4"}

	newSlice := StringSliceToInt(slice, 0)

	assert.EqualValues(t, []int{1, 2, 3, 4}, newSlice)
}

func Test_StringSliceToInt_bitOverflowSigned(t *testing.T) {
	slice := []string{"-200", "-100", "0", "100", "200"}

	newSlice := StringSliceToInt[int8](slice, 42)

	assert.EqualValues(t, []int8{42, -100, 0, 100, 42}, newSlice)
}

func Test_StringSliceToInt_bitOverflowUnsigned(t *testing.T) {
	slice := []string{"-100", "0", "100", "200", "300"}

	newSlice := StringSliceToInt[uint8](slice, 42)

	assert.EqualValues(t, []uint8{42, 0, 100, 200, 42}, newSlice)
}
