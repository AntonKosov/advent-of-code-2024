package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (orderingRules [][2]int, updates [][]int) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	for delimiter := 0; ; delimiter++ {
		if lines[delimiter] != "" {
			continue
		}

		orderingRules = slice.Map(lines[:delimiter], func(line string) [2]int {
			parts := transform.StrToInts(line)
			return [2]int{parts[0], parts[1]}
		})

		updates = slice.Map(lines[delimiter+1:], func(line string) []int {
			return transform.StrToInts(line)
		})

		return orderingRules, updates
	}
}

func process(orderingRules [][2]int, updates [][]int) int {
	sum := 0
	lessTable := buildLessTable(orderingRules)
	for _, update := range updates {
		if correctUpdate(update, lessTable) {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func correctUpdate(update []int, lessTable LessTable) bool {
	for i, leftPageNumber := range update {
		less := lessTable[leftPageNumber]
		for j := i + 1; j < len(update); j++ {
			rightPageNumber := update[j]
			if c, ok := less[rightPageNumber]; ok && !c {
				return false
			}
			if more := lessTable[rightPageNumber]; more != nil {
				if c, ok := more[leftPageNumber]; ok && c {
					return false
				}
			}
		}
	}

	return true
}

func buildLessTable(orderingRules [][2]int) LessTable {
	lessTable := LessTable{}
	for _, rule := range orderingRules {
		less := rule[0]
		if _, ok := lessTable[less]; !ok {
			lessTable[less] = map[int]bool{}
		}
		lessTable[less][rule[1]] = true
	}

	return lessTable
}

type LessTable map[int]map[int]bool
