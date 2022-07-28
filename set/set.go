package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}
