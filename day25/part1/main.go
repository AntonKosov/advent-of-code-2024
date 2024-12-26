package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
)

const (
	height = 7
	width  = 5
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []Schematic {
	lines := input.Lines()
	var schematics []Schematic
	for len(lines) > 0 {
		schematics = append(schematics, NewSchematic(lines[:height]))
		lines = lines[height+1:]
	}

	return schematics
}

func process(schematics []Schematic) int {
	locks, keys := decode(schematics)
	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			if !overlap(lock, key) {
				count++
			}
		}
	}

	return count
}

func overlap(lock, key Code) bool {
	for c := range width {
		if lock[c]+key[c] > height-2 {
			return true
		}
	}

	return false
}

func decode(schematics []Schematic) (locks, keys []Code) {
	for _, scematic := range schematics {
		target := &keys
		if scematic[0][0] == '#' {
			target = &locks
		}
		*target = append(*target, NewCode(scematic))
	}

	return
}

type Code [width]int

func NewCode(schematic Schematic) Code {
	var code Code
	for r := 1; r < len(schematic)-1; r++ {
		row := schematic[r]
		for c, v := range row {
			if v == '#' {
				code[c]++
			}
		}
	}

	return code
}

type Schematic [height][width]rune

func NewSchematic(lines []string) Schematic {
	var schematic Schematic
	for r, row := range lines {
		for c, v := range row {
			schematic[r][c] = v
		}
	}

	return schematic
}
