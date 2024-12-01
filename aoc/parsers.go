package aoc

import (
	"regexp"
	"strconv"
)

func StrToInt(str string) int {
	r, err := strconv.Atoi(str)
	Must(err)

	return r
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
