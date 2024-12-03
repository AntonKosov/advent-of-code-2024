package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
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
	right := map[int]int{}
	for _, v := range lists[1] {
		right[v]++
	}

	score := 0
	for _, v := range lists[0] {
		score += v * right[v]
	}

	return score
}
