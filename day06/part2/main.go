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
	startState := findStartState(territory)
	visited, _ := guardPath(territory, startState)
	obstructions := 0
	for pos := range visited {
		if pos == startState.pos {
			continue
		}
		testObstruction := &territory[pos.Y][pos.X]
		*testObstruction = obstruction
		if _, loop := guardPath(territory, startState); loop {
			obstructions++
		}
		*testObstruction = '.'
	}

	return obstructions
}

func guardPath(territory [][]rune, state State) (visited map[math.Vector2[int]]bool, loop bool) {
	states := make(map[State]bool, len(territory)*len(territory[0]))
	visited = make(map[math.Vector2[int]]bool, len(states))
	for {
		states[state] = true
		fwd := state.pos.Add(state.dir)
		if fwd.Y < 0 || fwd.X < 0 || fwd.Y >= len(territory) || fwd.X >= len(territory[fwd.Y]) {
			return visited, false
		}
		if territory[fwd.Y][fwd.X] == obstruction {
			state.dir = state.dir.RotateRight()
			continue
		}
		state.pos = fwd
		if states[state] {
			return visited, true
		}
		visited[fwd] = true
	}
}

func findStartState(territory [][]rune) State {
	for y, row := range territory {
		for x, v := range row {
			if v == '^' {
				return State{pos: math.NewVector2(x, y), dir: math.NewVector2(0, -1)}
			}
		}
	}

	panic("incorrect input")
}

type State struct {
	pos math.Vector2[int]
	dir math.Vector2[int]
}

const obstruction = '#'
