package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
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
	for y := range puzzle {
		for x := range puzzle {
			for _, dir := range directions {
				if xmas(puzzle, math.NewVector2(x, y), dir) {
					count++
				}
			}
		}
	}

	return count
}

func xmas(puzzle [][]rune, pos, dir math.Vector2[int]) bool {
	for _, char := range "XMAS" {
		if pos.X < 0 || pos.Y < 0 || pos.Y >= len(puzzle) || pos.X >= len(puzzle[pos.Y]) {
			return false
		}
		if puzzle[pos.Y][pos.X] != char {
			return false
		}
		pos = pos.Add(dir)
	}

	return true
}

var directions = []math.Vector2[int]{
	{X: 0, Y: -1},
	{X: 1, Y: -1},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 0, Y: 1},
	{X: -1, Y: 1},
	{X: -1, Y: 0},
	{X: -1, Y: -1},
}
