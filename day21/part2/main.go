package main

import (
	"fmt"
	smath "math"
	"slices"
	"strings"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

const robotsInBetween = 25

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []string {
	lines := input.Lines()
	return lines[:len(lines)-1]
}

func process(codes []string) uint64 {
	complexity := uint64(0)
	for _, code := range codes {
		num := transform.StrToUInt64(code[:3])
		length := finalKeypadSequenceLength(code)
		complexity += num * length
	}

	return complexity
}

func finalKeypadSequenceLength(code string) uint64 {
	keypads := []Keypad{NumericKeypad}
	keypads = append(keypads, slices.Repeat([]Keypad{DirectionalKeypad}, robotsInBetween)...)
	return minSeqLength(code, keypads, Cache{})
}

func minSeqLength(seq string, keypads []Keypad, cache Cache) uint64 {
	minSeq := uint64(0)
	runes := []rune(string(ActOp) + seq)
	for i := 0; i < len(runes)-1; i++ {
		minSeq += minButtonPressSequence(runes[i], runes[i+1], keypads, cache)
	}

	return minSeq
}

func minButtonPressSequence(currentButton, targetButton rune, keypads []Keypad, cache Cache) uint64 {
	movement := Movement{from: currentButton, to: targetButton, depth: len(keypads)}
	if v, ok := cache[movement]; ok {
		return v
	}

	minSeq := uint64(smath.MaxUint64)
	for _, seq := range keypads[0](currentButton, targetButton) {
		if len(keypads) == 1 {
			minSeq = min(minSeq, uint64(len(seq)))
		} else {
			minSeq = min(minSeq, minSeqLength(seq, keypads[1:], cache))
		}
	}

	cache[movement] = minSeq

	return minSeq
}

type Keypad func(from, to rune) []string

var NumericKeypad Keypad

var DirectionalKeypad Keypad

func newKeypad(layout [][]rune) Keypad {
	positions := map[rune]math.Vector2[int]{}
	for y, row := range layout {
		for x, b := range row {
			positions[b] = math.NewVector2(x, y)
		}
	}

	type cacheKey struct{ from, to rune }
	cache := map[cacheKey][]string{}

	return func(from, to rune) (sequences []string) {
		key := cacheKey{from: from, to: to}
		if s, ok := cache[key]; ok {
			return s
		}

		defer func() { cache[key] = sequences }()

		if from == to {
			return []string{string(ActOp)}
		}

		currentPos := positions[from]
		targetPos := positions[to]

		straight := func(offset math.Vector2[int]) string {
			var seq strings.Builder
			dir := offset.Norm()
			for offset.X != 0 || offset.Y != 0 {
				seq.WriteRune(keyDir[dir])
				offset = offset.Sub(dir)
			}
			return seq.String()
		}

		offset := targetPos.Sub(currentPos)
		if offset.X == 0 || offset.Y == 0 {
			return []string{straight(offset) + string(ActOp)}
		}

		rotations := []math.Vector2[int]{
			math.NewVector2(currentPos.X, targetPos.Y),
			math.NewVector2(targetPos.X, currentPos.Y),
		}

		for _, rot := range rotations {
			if rot == positions[NoOp] {
				continue
			}
			seq1 := straight(rot.Sub(currentPos))
			seq2 := straight(targetPos.Sub(rot))
			sequences = append(sequences, seq1+seq2+string(ActOp))
		}

		return sequences
	}
}

func init() {
	NumericKeypad = newKeypad([][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{NoOp, '0', ActOp},
	})
	DirectionalKeypad = newKeypad([][]rune{
		{NoOp, '^', ActOp},
		{'<', 'v', '>'},
	})
}

var keyDir = map[math.Vector2[int]]rune{
	math.NewVector2(0, -1): '^',
	math.NewVector2(0, 1):  'v',
	math.NewVector2(1, 0):  '>',
	math.NewVector2(-1, 0): '<',
}

const (
	NoOp  = ' '
	ActOp = 'A'
)

type Movement struct {
	from  rune
	to    rune
	depth int
}

type Cache map[Movement]uint64
