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

func Test_IntSliceToStrings(t *testing.T) {
	slice := New(1, 2, 3, 4)

	newSlice := IntSliceToStrings(slice)

	assert.EqualValues(t, []string{"1", "2", "3", "4"}, newSlice)
}

func Test_UintSliceToStrings(t *testing.T) {
	slice := []uint32{1, 2, 3, 4}

	newSlice := UintSliceToStrings(slice)

	assert.EqualValues(t, []string{"1", "2", "3", "4"}, newSlice)
}
