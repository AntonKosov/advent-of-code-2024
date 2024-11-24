package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []string {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	return lines
}

func process(_ []string) int {
	return 0
}
