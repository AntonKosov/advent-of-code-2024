package transform

import (
	"regexp"
	"strconv"

	"github.com/AntonKosov/advent-of-code-2024/aoc/must"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
)

func StrToInt(str string) int {
	return must.Return(strconv.Atoi(str))
}

func StrToInts(str string) []int {
	return slice.Map(parseNums(str), func(num string) int { return StrToInt(num) })
}

func StrToUInt64(str string) uint64 {
	return must.Return(strconv.ParseUint(str, 10, 64))
}

func StrToUint64s(str string) []uint64 {
	return slice.Map(parseNums(str), func(num string) uint64 { return StrToUInt64(num) })
}

func parseNums(str string) []string {
	return regexp.MustCompile(`-?[\d]+`).FindAllString(str, -1)
}
