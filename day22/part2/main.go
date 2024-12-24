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

func read() []uint64 {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) uint64 { return transform.StrToUInt64(line) })
}

func process(numbers []uint64) uint64 {
	sequences := map[Sequence]uint64{}
	for _, num := range numbers {
		prices := generatePrices(num)
		seq := makeSequences(prices)
		mergeSequences(sequences, seq)
	}

	maxBananas := uint64(0)
	for _, bananas := range sequences {
		maxBananas = max(maxBananas, bananas)
	}

	return maxBananas
}

func mergeSequences(target map[Sequence]uint64, source map[Sequence]int8) {
	for seq, price := range source {
		target[seq] += uint64(price)
	}
}

func makeSequences(prices []int8) map[Sequence]int8 {
	sequences := map[Sequence]int8{}
	diffs := makeDiffs(prices)
	for i := 0; i < len(diffs)-3; i++ {
		seq := Sequence{diffs[i], diffs[i+1], diffs[i+2], diffs[i+3]}
		if _, ok := sequences[seq]; !ok {
			sequences[seq] = prices[i+3]
		}
	}

	return sequences
}

func makeDiffs(prices []int8) []int8 {
	diffs := make([]int8, len(prices))
	diffs[0] = prices[0]
	for i := 1; i < len(diffs); i++ {
		diffs[i] = prices[i] - prices[i-1]
	}

	return diffs
}

func generatePrices(number uint64) []int8 {
	const count = 2000
	calcPrice := func() int8 { return int8(number % 10) }

	prices := make([]int8, count)
	prices[0] = calcPrice()

	generators := []func() uint64{
		func() uint64 { return number * 64 },
		func() uint64 { return number / 32 },
		func() uint64 { return number * 2048 },
	}

	for i := 1; i < count; i++ {
		for _, gen := range generators {
			number = (gen() ^ number) % 16777216
		}
		prices[i] = calcPrice()
	}

	return prices
}

type Sequence [4]int8
