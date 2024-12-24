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

func Sign[T Numbers](v T) int {
	if v == 0 {
		return 0
	}

	if v < 0 {
		return -1
	}

	return 1
}
