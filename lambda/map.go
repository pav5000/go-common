package lambda

func MapSlice[From, To any](in []From, fn func(From) To) []To {
	if in == nil {
		return nil
	}
	to := make([]To, len(in))
	for i := range in {
		to[i] = fn(in[i])
	}

	return to
}

func SliceToMap[Input, Value any, Key comparable](in []Input, fn func(Input) (Key, Value)) map[Key]Value {
	to := make(map[Key]Value, len(in))
	for _, item := range in {
		key, value := fn(item)
		to[key] = value
	}

	return to
}

func MapToSlice[Value, Output any, Key comparable](in map[Key]Value, fn func(Key, Value) Output) []Output {
	out := make([]Output, 0, len(in))
	for key, value := range in {
		out = append(out, fn(key, value))
	}

	return out
}
