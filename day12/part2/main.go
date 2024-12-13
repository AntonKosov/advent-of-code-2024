package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/day12/part2/solution"
)

func main() {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	answer := solution.Price(solution.Parse(lines))
	fmt.Printf("Answer: %v\n", answer)
}
