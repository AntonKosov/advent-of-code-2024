package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []ClawMachine {
	const offset = 10000000000000
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	machines := make([]ClawMachine, 0, len(lines)/4)
	for i := 0; i < len(lines); i += 4 {
		a := transform.StrToInts(lines[i])
		b := transform.StrToInts(lines[i+1])
		p := transform.StrToInts(lines[i+2])
		machines = append(machines, ClawMachine{
			PushA: math.NewVector2(a[0], a[1]),
			PushB: math.NewVector2(b[0], b[1]),
			Prize: math.NewVector2(p[0]+offset, p[1]+offset),
		})
	}

	return machines
}

func process(machines []ClawMachine) int64 {
	var tokens int64
	for _, machine := range machines {
		if minTokens := machine.MinTokens(); minTokens != nil {
			tokens += *minTokens
		}
	}

	return tokens
}

type ClawMachine struct {
	PushA math.Vector2[int]
	PushB math.Vector2[int]
	Prize math.Vector2[int]
}

func (m ClawMachine) MinTokens() *int64 {
	// System of liner equations:
	// | a * PushA.X + b * PushB.X = Prize.X
	// | a * PushA.Y + b * PushB.Y = Prize.Y
	bNumerator := m.PushA.X*m.Prize.Y - m.Prize.X*m.PushA.Y
	bDenominator := m.PushA.X*m.PushB.Y - m.PushB.X*m.PushA.Y
	if bNumerator%bDenominator != 0 {
		return nil
	}

	b := bNumerator / bDenominator
	aNumberator := m.Prize.X - b*m.PushB.X
	aDenominator := m.PushA.X
	if aNumberator%aDenominator != 0 {
		return nil
	}

	a := aNumberator / aDenominator

	tokens := int64(3*a + b)

	return &tokens
}
