package math

func CountDigits[T Numbers](num T) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}

	return count
}
