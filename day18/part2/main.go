package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

const (
	size = 71
	// size = 7
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

func process(corruptedBytes []math.Vector2[int]) string {
	grid := makeGrid()
	var path map[math.Vector2[int]]bool
	for _, cb := range corruptedBytes {
		grid[cb.Y][cb.X] = true
		if path != nil && !path[cb] {
			continue
		}
		if path = findPath(grid); path == nil {
			return fmt.Sprintf("%v,%v", cb.X, cb.Y)
		}
	}

	panic("not found")
}

func findPath(grid [][]bool) map[math.Vector2[int]]bool {
	target := math.NewVector2(size-1, size-1)
	dirs := []math.Vector2[int]{{X: 1}, {X: -1}, {Y: -1}, {Y: 1}}
	visited := makeGrid()
	pq := slice.NewPriorityQueue(func(a, b State) bool {
		return a.pos.ManhattanDst(target) < b.pos.ManhattanDst(target)
	})
	pq.Push(State{pos: math.NewVector2(0, 0)})

	for !pq.Empty() {
		state := pq.Pop()
		pos := state.pos
		if pos == target {
			return state.path()
		}
		if visited[pos.Y][pos.X] {
			continue
		}
		visited[pos.Y][pos.X] = true
		for _, dir := range dirs {
			np := pos.Add(dir)
			if np.Y < 0 || np.X < 0 || np.Y >= size || np.X >= size || grid[np.Y][np.X] {
				continue
			}
			pq.Push(State{pos: np, prevState: &state})
		}
	}

	return nil
}

type State struct {
	pos       math.Vector2[int]
	prevState *State
}

func (s State) path() map[math.Vector2[int]]bool {
	path := map[math.Vector2[int]]bool{}
	for cs := &s; cs != nil; cs = cs.prevState {
		path[cs.pos] = true
	}

	return path
}

func makeGrid() [][]bool {
	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}

	return grid
}
