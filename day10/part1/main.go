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

func read() [][]byte {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) []byte {
		return slice.Map([]byte(line), func(v byte) byte { return v - '0' })
	})
}

func process(area [][]byte) int {
	sum := 0
	for y, row := range area {
		for x := range row {
			sum += countTrails(area, x, y)
		}
	}

	return sum
}

func countTrails(area [][]byte, x, y int) int {
	if area[y][x] != 0 {
		return 0
	}

	places := map[math.Vector2[int]]bool{math.NewVector2(x, y): true}
	for step := byte(1); step <= 9; step++ {
		nextSteps := map[math.Vector2[int]]bool{}
		for currentPos := range places {
			for _, dir := range dirs {
				pos := currentPos.Add(dir)
				if pos.X < 0 || pos.Y < 0 || pos.Y >= len(area) || pos.X >= len(area[pos.Y]) {
					continue
				}
				if area[pos.Y][pos.X] == step {
					nextSteps[pos] = true
				}
			}

		}

		places = nextSteps
	}

	return len(places)
}

var dirs = []math.Vector2[int]{{X: 1}, {X: -1}, {Y: 1}, {Y: -1}}
