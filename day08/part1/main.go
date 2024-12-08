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

func process(grid [][]rune) int {
	locations := collectLocations(grid)
	width, height := len(grid[0]), len(grid)
	antinodes := map[math.Vector2[int]]bool{}
	addAntinode := func(from, to math.Vector2[int]) {
		if pos := to.Add(to.Sub(from)); pos.X >= 0 && pos.X < width && pos.Y >= 0 && pos.Y < height {
			antinodes[pos] = true
		}
	}
	for _, antennas := range locations {
		for i := 0; i < len(antennas)-1; i++ {
			a1 := antennas[i]
			for j := i + 1; j < len(antennas); j++ {
				a2 := antennas[j]
				addAntinode(a1, a2)
				addAntinode(a2, a1)
			}
		}
	}

	return len(antinodes)
}

func collectLocations(grid [][]rune) (locations map[rune][]math.Vector2[int]) {
	locations = map[rune][]math.Vector2[int]{}
	for r, row := range grid {
		for c, v := range row {
			if v != '.' {
				locations[v] = append(locations[v], math.NewVector2(c, r))
			}
		}
	}

	return locations
}
