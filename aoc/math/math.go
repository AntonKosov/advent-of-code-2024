package math

type Numbers interface {
	int
}

func Abs[T Numbers](v T) T {
	if v < 0 {
		return -v
	}

	return v
}
