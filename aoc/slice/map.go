package slice

func Map[S, T any](source []S, transform func(S) T) []T {
	target := make([]T, len(source))
	for i, s := range source {
		target[i] = transform(s)
	}

	return target
}
