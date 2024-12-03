package main

import (
	"fmt"
	"regexp"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/must"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

func main() {
	data := string(input.Raw())
	mulRegex := must.Return(regexp.Compile(`mul\(\d+,\d+\)`))
	sum := 0
	for _, m := range mulRegex.FindAllString(data, -1) {
		parts := transform.StrToInts(m)
		sum += parts[0] * parts[1]
	}

	fmt.Printf("Answer: %v\n", sum)
}
