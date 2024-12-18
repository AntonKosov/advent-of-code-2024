package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

type Register rune

const (
	RegisterA Register = 'A'
	RegisterB Register = 'B'
	RegisterC Register = 'C'
)

type RegisterValues map[Register]int

type Instruction func(registers RegisterValues, output *[]int, ptr *int)

type Computer struct {
	instructionPtr int
	registers      RegisterValues
	instructions   []Instruction
	output         []int
}

func (c *Computer) Run() {
	for c.instructionPtr < len(c.instructions) {
		c.instructions[c.instructionPtr](c.registers, &c.output, &c.instructionPtr)
	}
}

func read() Computer {
	lines := input.Lines()

	computer := Computer{
		registers: RegisterValues{
			RegisterA: transform.StrToInts(lines[0])[0],
			RegisterB: transform.StrToInts(lines[1])[0],
			RegisterC: transform.StrToInts(lines[2])[0],
		},
	}

	comboValue := func(registers RegisterValues, code int) int {
		if code < 4 {
			return code
		}
		if code == 4 {
			return registers[RegisterA]
		}
		if code == 5 {
			return registers[RegisterB]
		}
		if code == 6 {
			return registers[RegisterC]
		}
		panic("unsupported combo")
	}

	program := transform.StrToInts(lines[4])
	for i := 0; i < len(program); i += 2 {
		opcode, operand := program[i], program[i+1]
		var instruction Instruction
		switch opcode {
		case 0: // adv
			instruction = func(registers RegisterValues, _ *[]int, ptr *int) {
				registers[RegisterA] = registers[RegisterA] / math.Pow(2, uint(comboValue(registers, operand)))
				*ptr++
			}
		case 1: // bxl
			instruction = func(registers RegisterValues, _ *[]int, ptr *int) {
				registers[RegisterB] = registers[RegisterB] ^ operand
				*ptr++
			}
		case 2: // bst
			instruction = func(registers RegisterValues, _ *[]int, ptr *int) {
				registers[RegisterB] = comboValue(registers, operand) & 0b111
				*ptr++
			}
		case 3: // jnz
			instruction = func(registers RegisterValues, _ *[]int, ptr *int) {
				if a := registers[RegisterA]; a != 0 {
					*ptr = operand / 2
					return
				}
				*ptr++
			}
		case 4: // bxc
			instruction = func(registers RegisterValues, _ *[]int, ptr *int) {
				registers[RegisterB] = registers[RegisterB] ^ registers[RegisterC]
				*ptr++
			}
		case 5: // out
			instruction = func(registers RegisterValues, output *[]int, ptr *int) {
				*output = append(*output, comboValue(registers, operand)&0b111)
				*ptr++
			}
		case 6: // bdv
			instruction = func(registers RegisterValues, _ *[]int, ptr *int) {
				registers[RegisterB] = registers[RegisterA] / math.Pow(2, uint(comboValue(registers, operand)))
				*ptr++
			}
		case 7: // cdv
			instruction = func(registers RegisterValues, _ *[]int, ptr *int) {
				registers[RegisterC] = registers[RegisterA] / math.Pow(2, uint(comboValue(registers, operand)))
				*ptr++
			}
		default:
			panic("unknown opcode")
		}
		computer.instructions = append(computer.instructions, instruction)
	}

	return computer
}

func process(computer Computer) string {
	computer.Run()

	var sb strings.Builder
	for _, v := range computer.output {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(strconv.Itoa(v))
	}

	return sb.String()
}
