package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []string {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return lines
}

func process(_ []string) int {
	return 0
}
