package main

import (
	"fmt"
	"sort"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [2][]int {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	input := [2][]int{make([]int, len(lines)), make([]int, len(lines))}
	for i, line := range lines {
		parts := transform.StrToInts(line)
		input[0][i] = parts[0]
		input[1][i] = parts[1]
	}

	return input
}

func process(lists [2][]int) int {
	sort.Ints(lists[0])
	sort.Ints(lists[1])

	distance := 0
	for i, v0 := range lists[0] {
		distance += math.Abs(v0 - lists[1][i])
	}

	return distance
}
