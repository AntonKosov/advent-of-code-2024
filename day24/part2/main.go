package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (vars map[string]uint8, commands map[Command]bool) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	vars = map[string]uint8{}
	for i, line := range lines {
		if line == "" {
			lines = lines[i+1:]
			break
		}
		vars[line[:3]] = uint8(line[5] - '0')
	}

	commands = make(map[Command]bool, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		command := Command{
			var1:      parts[0],
			logic:     parts[1],
			var2:      parts[2],
			targetVar: parts[4],
		}
		commands[command] = true
	}

	return vars, commands
}

func process(vars map[string]uint8, commands map[Command]bool) string {
	for name := range vars {
		vars[name] = 0
	}
	for i := 0; i < 45; i++ {
		varsCopy := maps.Clone(vars)
		x, y := fmt.Sprintf("x%02d", i), fmt.Sprintf("y%02d", i)
		varsCopy[x] = 1
		varsCopy[y] = 1
		s := sum(varsCopy, maps.Clone(commands))
		if s&(1<<i) != 0 {
			fmt.Printf("Bit #%02d: incorrect z%02d\n", i, i)
		}
		if s&(1<<(i+1)) == 0 {
			fmt.Printf("Bit #%02d: incorrect z%02d\n", i, i+1)
		}
	}

	/*
		Incorrect wire connections should be identified manually based on the incorrect results.

		Example:

		bct OR ggp -> chn	// z39: chn is the carry of 40th bit

		y40 AND x40 -> rcb	// Carry
		x40 XOR y40 -> gfd	// Sum
		chn AND gfd -> prp	// Sum up carries
		rcb OR prp -> bbr	// bbr is the carry of 41st bit
		chn XOR gfd -> z40	// Result
	*/

	// x21 AND y21 -> z21 <---> nsp XOR tqh -> gds
	// snp OR mnh -> z15 <---> ccp XOR hhw -> fph
	// y30 AND x30 -> wrk <---> y30 XOR x30 -> jrs
	// ksm XOR fcv -> cqk <---> ksm AND fcv -> z34
	v := []string{"z21", "gds", "z15", "fph", "wrk", "jrs", "cqk", "z34"}
	slices.Sort(v)

	return strings.Join(v, ",")
}

func sum(vars map[string]uint8, commands map[Command]bool) uint64 {
	for len(commands) > 0 {
		var usedCommands []Command
		for command := range commands {
			var v1, v2 uint8
			var ok bool
			if v1, ok = vars[command.var1]; !ok {
				continue
			}
			if v2, ok = vars[command.var2]; !ok {
				continue
			}
			usedCommands = append(usedCommands, command)
			var output uint8
			switch command.logic {
			case "AND":
				output = v1 & v2
			case "OR":
				output = v1 | v2
			case "XOR":
				output = v1 ^ v2
			default:
				panic("unknown command")
			}
			vars[command.targetVar] = output
		}

		for _, usedCommand := range usedCommands {
			delete(commands, usedCommand)
		}
	}

	var result uint64
	for name, value := range vars {
		if name[0] != 'z' || value == 0 {
			continue
		}
		bit := transform.StrToInt(name[1:])
		result = result | (1 << bit)
	}

	return result
}

type Command struct {
	var1      string
	var2      string
	logic     string
	targetVar string
}
