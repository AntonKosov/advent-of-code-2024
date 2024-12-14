package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

const (
	// width  = 11
	// height = 7
	width  = 101
	height = 103
	moves  = 100
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []Robot {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) Robot {
		values := transform.StrToInts(line)
		return Robot{
			position: math.NewVector2(values[0], values[1]),
			velocity: math.NewVector2(values[2], values[3]),
		}
	})
}

func process(robots []Robot) int {
	for i := range robots {
		r := &robots[i]
		pos := r.position.Add(r.velocity.Mul(moves))
		r.position = math.NewVector2(math.Mod(pos.X, width), math.Mod(pos.Y, height))
	}

	return calcSafety(robots)
}

func calcSafety(robots []Robot) int {
	var quadrants [2][2]int
	cw, ch := width/2, height/2
	for _, robot := range robots {
		pos := robot.position
		if pos.X == cw || pos.Y == ch {
			continue
		}
		qx, qy := 0, 0
		if pos.X > cw {
			qx++
		}
		if pos.Y > ch {
			qy++
		}
		quadrants[qy][qx]++
	}

	return quadrants[0][0] * quadrants[0][1] * quadrants[1][0] * quadrants[1][1]
}

type Robot struct {
	position math.Vector2[int]
	velocity math.Vector2[int]
}
