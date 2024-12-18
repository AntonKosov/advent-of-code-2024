package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

const (
	size        = 71
	fallenBytes = 1024
	// size        = 7
	// fallenBytes = 12
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []math.Vector2[int] {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) math.Vector2[int] {
		nums := transform.StrToInts(line)
		return math.NewVector2(nums[0], nums[1])
	})
}

func process(corruptedBytes []math.Vector2[int]) int {
	grid := makeGrid(corruptedBytes)
	dirs := []math.Vector2[int]{{X: 1}, {X: -1}, {Y: -1}, {Y: 1}}
	target := math.NewVector2(size-1, size-1)
	visited := map[math.Vector2[int]]bool{}
	pq := slice.NewPriorityQueue(func(a, b State) bool {
		return a.moves+a.pos.ManhattanDst(target) < b.moves+b.pos.ManhattanDst(target)
	})
	pq.Push(State{pos: math.NewVector2(0, 0)})

	for !pq.Empty() {
		state := pq.Pop()
		pos := state.pos
		if pos == target {
			return state.moves
		}
		if visited[pos] {
			continue
		}
		visited[pos] = true
		for _, dir := range dirs {
			np := pos.Add(dir)
			if np.Y < 0 || np.X < 0 || np.Y >= size || np.X >= size || grid[np.Y][np.X] {
				continue
			}
			pq.Push(State{pos: np, moves: state.moves + 1})
		}
	}

	panic("not found")
}

type State struct {
	pos   math.Vector2[int]
	moves int
}

func makeGrid(corruptedBytes []math.Vector2[int]) [][]bool {
	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}

	for i := 0; i < fallenBytes; i++ {
		cb := corruptedBytes[i]
		grid[cb.Y][cb.X] = true
	}

	return grid
}
