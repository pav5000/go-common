package set

// Set is a collection of unique elements.
type Set[T comparable] map[T]struct{}

// New creates empty set.
func New[T comparable]() Set[T] {
	return make(Set[T])
}

// WithCapacity creates a set with specified capacity.
func WithCapacity[T comparable](capacity int) Set[T] {
	return make(Set[T], capacity)
}

// WithCapacity creates a set with elements of slice.
func FromSlice[T comparable](slice []T) Set[T] {
	set := make(Set[T], len(slice))
	for _, elem := range slice {
		set[elem] = struct{}{}
	}

	return set
}

// Add adds the element into the set.
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// Remove removes the element from the set.
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// Has checks if the set contains the element.
func (s Set[T]) Has(elem T) bool {
	_, ok := s[elem]

	return ok
}

// AsSlice converts the set to a slice.
func (s Set[T]) AsSlice() []T {
	slice := make([]T, 0, len(s))
	for elem := range s {
		slice = append(slice, elem)
	}

	return slice
}

// Copy creates the new set from the same elements.
func (s Set[T]) Copy() Set[T] {
	set := make(Set[T], len(s))
	for elem := range s {
		set[elem] = struct{}{}
	}

	return set
}

// Filter removes elements for which fn returned false
// mutates the set inplace.
func (s Set[T]) Filter(fn func(T) bool) {
	for elem := range s {
		if !fn(elem) {
			delete(s, elem)
		}
	}
}

// Each calls fn for each element, stops when fn returns false.
func (s Set[T]) Each(fn func(T) bool) {
	for elem := range s {
		if !fn(elem) {
			return
		}
	}
}

// Map calls fn for each element of the input set
// fn results are added to the output set.
func Map[T, R comparable](set Set[T], fn func(T) R) Set[R] {
	out := make(Set[R], len(set))
	for elem := range set {
		out[fn(elem)] = struct{}{}
	}

	return out
}
