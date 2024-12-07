package math

func Pow[T Numbers](base T, power uint) T {
	if power == 0 {
		return 1
	}

	result := base
	for i := uint(1); i < power; i++ {
		result *= base
	}

	return result
}
