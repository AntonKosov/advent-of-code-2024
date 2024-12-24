package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []uint64 {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) uint64 { return transform.StrToUInt64(line) })
}

func process(numbers []uint64) uint64 {
	sum := uint64(0)
	for _, num := range numbers {
		sum += calc(num, 2000)
	}

	return sum
}

func calc(number uint64, n int) uint64 {
	generators := []func() uint64{
		func() uint64 { return number * 64 },
		func() uint64 { return number / 32 },
		func() uint64 { return number * 2048 },
	}

	for i := 0; i < n; i++ {
		for _, gen := range generators {
			num := gen()
			number = (num ^ number) % 16777216
		}
	}

	return number
}
