package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() map[uint64]uint64 {
	nums := transform.StrToUint64s(input.Lines()[0])
	stones := map[uint64]uint64{}
	for _, num := range nums {
		stones[num]++
	}

	return stones
}

func process(stones map[uint64]uint64) uint64 {
	for range 75 {
		stones = blink(stones)
	}

	return countStones(stones)
}

func countStones(stones map[uint64]uint64) uint64 {
	var sum uint64
	for _, count := range stones {
		sum += count
	}

	return sum
}

func blink(stones map[uint64]uint64) map[uint64]uint64 {
	next := make(map[uint64]uint64, len(stones)*2)
	for num, count := range stones {
		if num == 0 {
			next[1] += count
			continue
		}

		if d := math.CountDigits(num); d%2 == 0 {
			exp := uint64(math.Pow(10, uint(d/2)))
			next[num/exp] += count
			next[num%exp] += count
			continue
		}

		next[num*2024] += count
	}

	return next
}
