package main

import (
	"fmt"
)

func findSourceA(targetA uint64, program []uint64) (uint64, bool) {
	const delimiter = 8
	output := program[len(program)-1]
	for i := uint64(0); i < delimiter; i++ {
		a := targetA*delimiter + i
		//2,4
		b := a & 0b111
		//1,3
		b = b ^ 0b011
		//7,5
		c := a / (1 << b)
		//0,3
		// a = a / 8 // delimiter
		//1,5
		b = b ^ 0b101
		//4,4
		b = b ^ c
		//5,5
		code := b & 0b111
		if code == output {
			if len(program) == 1 {
				return a, true
			}
			if sourceA, found := findSourceA(a, program[:len(program)-1]); found {
				return sourceA, true
			}
		}
	}

	return 0, false
}

func main() {
	program := []uint64{2, 4, 1, 3, 7, 5, 0, 3, 1, 5, 4, 4, 5, 5, 3, 0}
	sourceA, found := findSourceA(0, program)
	if !found {
		fmt.Println("no solution")
		return
	}
	fmt.Println("Answer:", sourceA)
}
