package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][2]string {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) [2]string { return [2]string{line[:2], line[3:]} })
}

func process(connections [][2]string) int {
	graph := buildGraph(connections)
	sets := map[string]bool{}
	for node1, node1Connections := range graph {
		for node2 := range node1Connections {
			node2Connections := graph[node2]
			for node3 := range node2Connections {
				if node1 == node3 || !node1Connections[node3] || !startsWithT(node1, node2, node3) {
					continue
				}
				nodes := []string{node1, node2, node3}
				slices.Sort(nodes)
				sets[strings.Join(nodes, "")] = true
			}
		}
	}

	return len(sets)
}

func startsWithT(names ...string) bool {
	for _, name := range names {
		if name[0] == 't' {
			return true
		}
	}

	return false
}

func buildGraph(connections [][2]string) map[string]map[string]bool {
	graph := map[string]map[string]bool{}
	for _, connection := range connections {
		comp1, comp2 := connection[0], connection[1]
		m1, m2 := graph[comp1], graph[comp2]
		if m1 == nil {
			m1 = map[string]bool{}
			graph[comp1] = m1
		}
		if m2 == nil {
			m2 = map[string]bool{}
			graph[comp2] = m2
		}
		m1[comp2] = true
		m2[comp1] = true
	}

	return graph
}
