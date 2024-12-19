package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (patterns, designs []string) {
	lines := input.Lines()

	return strings.Split(lines[0], ", "), lines[2 : len(lines)-1]
}

func process(patterns, designs []string) int {
	tri := makeTri(patterns)
	count := 0
	for _, design := range designs {
		if possibleDesign([]rune(design), tri) {
			count++
		}
	}

	return count
}

func possibleDesign(design []rune, triRoot *Tri) bool {
	starts := make([]bool, len(design))
	starts[0] = true
	for i := range design {
		if !starts[i] {
			continue
		}
		for length := range designLengths(design[i:], triRoot) {
			idx := i + length
			if idx == len(design) {
				return true
			}
			starts[idx] = true
		}
	}

	return false
}

func designLengths(design []rune, triRoot *Tri) func(func(int) bool) {
	return func(yield func(int) bool) {
		node := triRoot
		for i, char := range design {
			node = node.chars[char]
			if node == nil {
				break
			}
			if node.last {
				if !yield(i + 1) {
					return
				}
			}
		}
	}
}

func makeTri(patterns []string) *Tri {
	root := NewTri()
	for _, pattern := range patterns {
		node := root
		for i, ch := range pattern {
			node = node.add(ch, i+1 == len(pattern))
		}
	}

	return root
}

type Tri struct {
	chars map[rune]*Tri
	last  bool
}

func NewTri() *Tri {
	return &Tri{chars: make(map[rune]*Tri, 5)}
}

func (t *Tri) add(char rune, last bool) *Tri {
	if tri, ok := t.chars[char]; ok {
		tri.last = tri.last || last
		return tri
	}

	tri := NewTri()
	tri.last = last
	t.chars[char] = tri

	return tri
}
