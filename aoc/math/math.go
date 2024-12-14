package math

type Numbers interface {
	int8 | int | uint64
}

func Abs[T Numbers](v T) T {
	if v < 0 {
		return -v
	}

	return v
}

func Mod(a, b int) int {
	return (a%b + b) % b
}
