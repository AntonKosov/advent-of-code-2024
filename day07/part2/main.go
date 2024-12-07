package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []Equation {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) Equation {
		parts := transform.StrToUint64s(line)
		return Equation{result: parts[0], numbers: parts[1:]}
	})
}

func process(equations []Equation) uint64 {
	calibration := uint64(0)
	for _, e := range equations {
		if trueEquation(e.result, e.numbers[0], e.numbers[1:]) {
			calibration += e.result
		}
	}

	return calibration
}

func trueEquation(result, value uint64, numbers []uint64) bool {
	if len(numbers) == 0 {
		return result == value
	}

	num, restNums := numbers[0], numbers[1:]
	combinedNum := value*math.Pow(uint64(10), uint(math.CountDigits(num))) + num

	return result >= value && (trueEquation(result, value*num, restNums) ||
		trueEquation(result, value+num, restNums) ||
		trueEquation(result, combinedNum, restNums))
}

type Equation struct {
	result  uint64
	numbers []uint64
}
