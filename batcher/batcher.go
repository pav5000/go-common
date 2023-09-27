package batcher

// Batch slices a slice into several batches
// each batch's length is less or equal than batchSize
// keep in mind that batches are slices of the base array and share the same memory
func Batch[T any](elems []T, batchSize int) [][]T {
	if len(elems) <= batchSize {
		return [][]T{elems}
	}
	divBatchCount := len(elems) / batchSize
	batches := make([][]T, 0, divBatchCount+1)
	divStop := divBatchCount * batchSize
	i := 0
	for i = 0; i < divStop; i += batchSize {
		batches = append(batches, elems[i:i+batchSize])
	}

	if i < len(elems) {
		batches = append(batches, elems[i:])
	}

	return batches
}
