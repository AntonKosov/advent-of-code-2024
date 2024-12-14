package main

import (
	"fmt"
	smath "math"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

const (
	width  = 101
	height = 103
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

type Tiles map[math.Vector2[int]][]math.Vector2[int]

func (t Tiles) print() {
	for r := range height {
		for c := range width {
			if c := t[math.NewVector2(c, r)]; c != nil {
				fmt.Print(len(c))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func read() Tiles {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	tiles := make(Tiles, len(lines))
	for _, line := range lines {
		values := transform.StrToInts(line)
		position := math.NewVector2(values[0], values[1])
		velocity := math.NewVector2(values[2], values[3])
		tiles[position] = append(tiles[position], velocity)
	}

	return tiles
}

type EntropyValues []int

func (v EntropyValues) entropy() float64 {
	var e float64
	for _, value := range v {
		if value == 0 {
			continue
		}
		ev := float64(value) / float64(len(v))
		e += ev * smath.Log2(ev)
	}

	return -e
}

// https://en.wikipedia.org/wiki/Entropy_(information_theory)
func process(tiles Tiles) int {
	horEntropy, verEntropy := make(EntropyValues, width), make(EntropyValues, height)
	for pos, vel := range tiles {
		horEntropy[pos.X] += len(vel)
		verEntropy[pos.Y] += len(vel)
	}
	minEntropy := horEntropy.entropy() + verEntropy.entropy()
	bestMove := 0
	for move := 1; move < bestMove+10000; move++ {
		nextTiles := make(Tiles, len(tiles)*2)
		for pos, velocities := range tiles {
			for _, velocity := range velocities {
				nextPos := pos.Add(velocity)
				nextPos = math.NewVector2(math.Mod(nextPos.X, width), math.Mod(nextPos.Y, height))
				nextTiles[nextPos] = append(nextTiles[nextPos], velocity)
				horEntropy[pos.X]--
				verEntropy[pos.Y]--
				horEntropy[nextPos.X]++
				verEntropy[nextPos.Y]++
			}
		}
		tiles = nextTiles
		if ent := horEntropy.entropy() + verEntropy.entropy(); ent < minEntropy {
			fmt.Println("===================================================================")
			fmt.Printf("moves: %v (+%v), entropy: %v (%v)\n", move, move-bestMove, ent, ent-minEntropy)
			tiles.print()
			bestMove = move
			minEntropy = ent
		}
	}

	return bestMove
}
