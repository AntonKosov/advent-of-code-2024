package aoc

import (
	"os"
	"strings"
)

func ReadRawInput() []byte {
	if len(os.Args) != 2 {
		panic("wrong arguments")
	}

	bytes, err := os.ReadFile(os.Args[1])
	Must(err)

	return bytes
}

func ReadAllInput() []string {
	return strings.Split(string(ReadRawInput()), "\n")
}
