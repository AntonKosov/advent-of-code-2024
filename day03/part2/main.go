package main

import (
	"fmt"
	"regexp"

	"github.com/AntonKosov/advent-of-code-2024/aoc"
)

func main() {
	data := string(aoc.ReadRawInput())
	mulRegex, err := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	aoc.Must(err)

	sum, enabled := 0, true
	for _, m := range mulRegex.FindAllString(data, -1) {
		if m[0] == 'd' {
			enabled = m == "do()"
		} else if enabled {
			parts := aoc.StrToInts(m)
			sum += parts[0] * parts[1]
		}
	}

	fmt.Printf("Answer: %v\n", sum)
}
