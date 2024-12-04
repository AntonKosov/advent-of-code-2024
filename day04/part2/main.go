package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]rune {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) []rune { return []rune(line) })
}

func process(puzzle [][]rune) int {
	count := 0
	for y := 1; y < len(puzzle)-1; y++ {
		for x := 1; x < len(puzzle[y])-1; x++ {
			if xmas(puzzle, x, y) {
				count++
			}
		}
	}

	return count
}

func xmas(puzzle [][]rune, x, y int) bool {
	if puzzle[y][x] != 'A' {
		return false
	}

	for _, dir := range [][2]int{{-1, 1}, {1, 1}} {
		dx, dy := dir[0], dir[1]
		c1, c2 := puzzle[y+dy][x+dx], puzzle[y-dy][x-dx]
		if !((c1 == 'M' && c2 == 'S') || (c1 == 'S' && c2 == 'M')) {
			return false
		}
	}

	return true
}
