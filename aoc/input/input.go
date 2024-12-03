package input

import (
	"os"
	"strings"

	"github.com/AntonKosov/advent-of-code-2024/aoc/must"
)

func Raw() []byte {
	if len(os.Args) != 2 {
		panic("wrong arguments")
	}

	return must.Return(os.ReadFile(os.Args[1]))
}

func Lines() []string {
	return strings.Split(string(Raw()), "\n")
}
