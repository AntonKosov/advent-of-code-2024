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

type State struct {
	pos, dir  math.Vector2[int]
	score     int
	prevState *State
}

type VisitedState struct {
	pos, dir math.Vector2[int]
}

func process(walls [][]bool, start, end math.Vector2[int]) int {
	pq := slice.NewPriorityQueue(func(s1, s2 State) bool {
		return s1.score+s1.pos.ManhattanDst(end) < s2.score+s2.pos.ManhattanDst(end)
	})
	visitedStates := map[VisitedState]int{}
	winningTiles := map[math.Vector2[int]]bool{}
	var winningScore *int
	pq.Push(State{pos: start, dir: math.NewVector2(1, 0)})
	for !pq.Empty() {
		state := pq.Pop()
		if state.pos == end {
			if winningScore != nil && *winningScore < state.score {
				continue
			}
			winningScore = &state.score
			for s := &state; s != nil; s = s.prevState {
				winningTiles[s.pos] = true
			}
			continue
		}

		push := func(pos, dir math.Vector2[int], scoreInc int) {
			if walls[pos.Y][pos.X] {
				return
			}
			nextState := State{
				pos:       pos,
				dir:       dir,
				score:     state.score + scoreInc,
				prevState: &state,
			}
			vs := VisitedState{pos: pos, dir: dir}
			if score, ok := visitedStates[vs]; ok && score < nextState.score {
				return
			}
			visitedStates[vs] = nextState.score
			pq.Push(nextState)
		}

		fwdDir := state.dir
		leftDir, rightDir := fwdDir.RotateLeft(), fwdDir.RotateRight()
		pos := state.pos
		push(pos.Add(fwdDir), fwdDir, 1)
		push(pos.Add(leftDir), leftDir, 1001)
		push(pos.Add(rightDir), rightDir, 1001)
	}

	return len(winningTiles)
}
