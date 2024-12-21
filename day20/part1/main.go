package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
)

const (
	// leastPico    = 1
	leastPico    = 100
	cheatingTime = 2
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (walls [][]bool, start, end math.Vector2[int]) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	walls = make([][]bool, len(lines))
	for y, line := range lines {
		walls[y] = make([]bool, len(line))
		for x, v := range line {
			switch v {
			case '#':
				walls[y][x] = true
			case 'S':
				start = math.NewVector2(x, y)
			case 'E':
				end = math.NewVector2(x, y)
			}
		}
	}

	return
}

func process(walls [][]bool, start, end math.Vector2[int]) int {
	stepsLeft := countSteps(walls, start, end)
	count := 0
	for pos, steps := range stepsLeft {
		for _, offset := range cheatOffsets {
			ce := pos.Add(offset)
			if ce.Y <= 0 || ce.X <= 0 || ce.Y >= len(walls)-1 || ce.X >= len(walls[0])-1 || walls[ce.Y][ce.X] {
				continue
			}
			if diff := steps - pos.ManhattanDst(ce) - stepsLeft[ce]; diff >= leastPico {
				count++
			}
		}
	}

	return count
}

func countSteps(walls [][]bool, start, end math.Vector2[int]) map[math.Vector2[int]]int {
	path := []math.Vector2[int]{start}
	dirs := []math.Vector2[int]{{X: -1}, {X: 1}, {Y: 1}, {Y: -1}}
	for prev, current := start, start; current != end; {
		for _, dir := range dirs {
			pos := current.Add(dir)
			if walls[pos.Y][pos.X] || pos == prev {
				continue
			}
			path = append(path, pos)
			prev, current = current, pos
			break
		}
	}

	cache := make(map[math.Vector2[int]]int, len(path))
	for i, pos := range path {
		cache[pos] = len(path) - i - 1
	}

	return cache
}

var cheatOffsets []math.Vector2[int]

func init() {
	origin := math.NewVector2(0, 0)
	for y := -cheatingTime; y <= cheatingTime; y++ {
		for x := -cheatingTime; x <= cheatingTime; x++ {
			offset := math.NewVector2(x, y)
			if dist := offset.ManhattanDst(origin); dist > 0 && dist <= cheatingTime {
				cheatOffsets = append(cheatOffsets, offset)
			}
		}
	}
}
