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
	antinodes := map[math.Vector2[int]]bool{}
	addAntinodes := func(from, to math.Vector2[int]) {
		for pos, offset := to, to.Sub(from); ; {
			pos = pos.Add(offset)
			if pos.X < 0 || pos.X >= len(grid[0]) || pos.Y < 0 || pos.Y >= len(grid) {
				return
			}
			antinodes[pos] = true
		}
	}
	for _, antennas := range locations {
		for _, antenna := range antennas {
			antinodes[antenna] = true
		}

		for i := 0; i < len(antennas)-1; i++ {
			a1 := antennas[i]
			for j := i + 1; j < len(antennas); j++ {
				a2 := antennas[j]
				addAntinodes(a1, a2)
				addAntinodes(a2, a1)
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
