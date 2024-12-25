package main

import (
	"fmt"
	"maps"
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

func process(connections [][2]string) string {
	largestSet := map[string]bool{}
	graph := buildGraph(connections)
nextNode:
	for node1, node1Connections := range graph {
		if len(node1Connections) <= len(largestSet) {
			continue
		}
		potentialSet := maps.Clone(node1Connections)
		exclude := map[string]bool{}
		for node2 := range potentialSet {
			node2Connections := graph[node2]
			for node3 := range potentialSet {
				if exclude[node3] {
					continue
				}

				if node2 != node3 && !node2Connections[node3] {
					exclude[node3] = true
					if len(potentialSet)-len(exclude)+1 <= len(largestSet) {
						continue nextNode
					}
				}
			}
			for ex := range exclude {
				delete(potentialSet, ex)
			}
			exclude = map[string]bool{}
		}
		largestSet = potentialSet
		largestSet[node1] = true
	}

	names := slices.Sorted(maps.Keys(largestSet))

	return strings.Join(names, ",")
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
