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
	mulRegex := must.Return(regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`))
	sum, enabled := 0, true
	for _, m := range mulRegex.FindAllString(data, -1) {
		if m[0] == 'd' {
			enabled = m == "do()"
		} else if enabled {
			parts := transform.StrToInts(m)
			sum += parts[0] * parts[1]
		}
	}

	fmt.Printf("Answer: %v\n", sum)
}
