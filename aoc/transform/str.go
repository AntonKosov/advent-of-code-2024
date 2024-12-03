package transform

import (
	"regexp"
	"strconv"

	"github.com/AntonKosov/advent-of-code-2024/aoc/must"
)

func StrToInt(str string) int {
	return must.Return(strconv.Atoi(str))
}

func StrToInts(str string) []int {
	r := regexp.MustCompile(`-?[\d]+`)
	matches := r.FindAllString(str, -1)

	res := make([]int, len(matches))
	for i, m := range matches {
		res[i] = StrToInt(m)
	}

	return res
}
