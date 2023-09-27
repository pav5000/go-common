package batcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Batch_WhenElemCountIsEqualToOneBatch_ShouldReturnOneBatch(t *testing.T) {
	batches := Batch([]int{1, 2, 3}, 3)

	assert.Equal(t, [][]int{{1, 2, 3}}, batches)
}

func Test_Batch_WhenElemCountIsLessThanOneBatch_ShouldReturnOneBatch(t *testing.T) {
	batches := Batch([]int{1, 2}, 3)

	assert.Equal(t, [][]int{{1, 2}}, batches)
}

func Test_Batch_WhenElemCountIsExactlyTripleTheBatchSize_ShouldReturnThreeBatches(t *testing.T) {
	batches := Batch([]int{1, 2, 3, 4, 5, 6}, 2)

	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}}, batches)
}

func Test_Batch_WhenElemCountIsNotDivisibleByBatchSize_TheLastBatchShouldCounainLessElements(t *testing.T) {
	batches := Batch([]int{1, 2, 3, 4, 5}, 3)

	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5}}, batches)
}

func Test_Batch_WhenElemCountIsNotDivisibleByBatchSize_TheLastBatchContainsOneElem(t *testing.T) {
	batches := Batch([]int{1, 2, 3, 4}, 3)

	assert.Equal(t, [][]int{{1, 2, 3}, {4}}, batches)
}
