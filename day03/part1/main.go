package main

import (
	"fmt"
	"regexp"

	"github.com/AntonKosov/advent-of-code-2024/aoc"
)

func main() {
	data := string(aoc.ReadRawInput())
	mulRegex, err := regexp.Compile(`mul\(\d+,\d+\)`)
	aoc.Must(err)

	sum := 0
	for _, m := range mulRegex.FindAllString(data, -1) {
		parts := aoc.StrToInts(m)
		sum += parts[0] * parts[1]
	}

	fmt.Printf("Answer: %v\n", sum)
}
