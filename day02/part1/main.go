package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]int {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) []int { return aoc.StrToInts(line) })
}

func process(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		if safeReport(report) {
			safeReports++
		}
	}

	return safeReports
}

func safeReport(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	sign := 1
	if report[0] > report[1] {
		sign = -1
	}

	for i := 0; i < len(report)-1; i++ {
		diff := (report[i+1] - report[i]) * sign
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
