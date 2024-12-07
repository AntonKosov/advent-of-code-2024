package math

type Numbers interface {
	int | uint64
}

func Abs[T Numbers](v T) T {
	if v < 0 {
		return -v
	}

	return v
}
