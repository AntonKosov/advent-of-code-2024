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

func read() [][]int8 {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) []int8 {
		return slice.Map([]rune(line), func(r rune) int8 { return int8(r) })
	})
}

func process(garden [][]int8) int {
	price := 0
	for y, row := range garden {
		for x := range row {
			area, perimeter := measure(garden, x, y)
			price += area * perimeter
		}
	}

	return price
}

func measure(garden [][]int8, x, y int) (area, perimeter int) {
	regionType := garden[y][x]
	plots := []math.Vector2[int]{math.NewVector2(x, y)}
	for len(plots) > 0 {
		plot := plots[len(plots)-1]
		plots = plots[:len(plots)-1]
		if garden[plot.Y][plot.X] < 0 {
			continue
		}

		garden[plot.Y][plot.X] *= -1
		area++
		for _, dir := range dirs {
			pos := plot.Add(dir)
			if pos.Y < 0 || pos.X < 0 || pos.Y >= len(garden) || pos.X >= len(garden[pos.Y]) {
				perimeter++
				continue
			}
			pt := garden[pos.Y][pos.X]
			if math.Abs(pt) != regionType {
				perimeter++
				continue
			}
			if pt == regionType {
				plots = append(plots, pos)
			}
		}
	}

	return area, perimeter
}

var dirs = []math.Vector2[int]{{X: -1}, {X: 1}, {Y: -1}, {Y: 1}}
