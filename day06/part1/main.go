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

func process(territory [][]rune) int {
	pos := findStartPosition(territory)
	dir := math.NewVector2(0, -1)
	visited := map[math.Vector2[int]]bool{}
	for {
		visited[pos] = true
		fwd := pos.Add(dir)
		if fwd.Y < 0 || fwd.X < 0 || fwd.Y >= len(territory) || fwd.X >= len(territory[fwd.Y]) {
			return len(visited)
		}
		if territory[fwd.Y][fwd.X] == '#' {
			dir = dir.RotateRight()
			continue
		}
		pos = fwd
	}
}

func findStartPosition(territory [][]rune) math.Vector2[int] {
	for y, row := range territory {
		for x, v := range row {
			if v == '^' {
				return math.NewVector2(x, y)
			}
		}
	}

	panic("incorrect input")
}
