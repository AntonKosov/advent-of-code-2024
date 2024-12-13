package solution_test

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2024/day12/part2/solution"
)

func runTest(t *testing.T, input []string, expectedPrice int) {
	garden := solution.Parse(input)
	actualPrice := solution.Price(garden)

	if actualPrice != expectedPrice {
		t.Logf("expected price: %v, actual price: %v\n", expectedPrice, actualPrice)
		t.Fail()
	}
}

func Test4(t *testing.T) {
	input := []string{
		"A",
	}
	runTest(t, input, 4)
}

func Test16(t *testing.T) {
	input := []string{
		"AA",
		"AA",
	}
	runTest(t, input, 16)
}

func Test80(t *testing.T) {
	input := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}
	runTest(t, input, 80)
}

func Test436(t *testing.T) {
	input := []string{
		"OOOOO",
		"OXOXO",
		"OOOOO",
		"OXOXO",
		"OOOOO",
	}
	runTest(t, input, 436)
}

func Test236(t *testing.T) {
	input := []string{
		"EEEEE",
		"EXXXX",
		"EEEEE",
		"EXXXX",
		"EEEEE",
	}
	runTest(t, input, 236)
}

func Test368(t *testing.T) {
	input := []string{
		"AAAAAA",
		"AAABBA",
		"AAABBA",
		"ABBAAA",
		"ABBAAA",
		"AAAAAA",
	}
	runTest(t, input, 368)
}

func Test1206(t *testing.T) {
	input := []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}
	runTest(t, input, 1206)
}

func Test88(t *testing.T) {
	input := []string{
		"AAAA",
		"ABCA",
		"AAAA",
	}
	runTest(t, input, 88)
}
