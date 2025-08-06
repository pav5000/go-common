package lambda

func FilterSliceInplace[S ~[]T, T any](s S, cb func(T) bool) S {
	filtered := s[:0]
	for _, item := range s {
		if cb(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func FilterSlice[S ~[]T, T any](s S, cb func(T) bool) S {
	filtered := make(S, 0)
	for _, item := range s {
		if cb(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func FilterMap[M ~map[K]V, K comparable, V any](m M, cb func(K, V) bool) M {
	filtered := make(M)
	for key, value := range m {
		if cb(key, value) {
			filtered[key] = value
		}
	}

	return filtered
}
