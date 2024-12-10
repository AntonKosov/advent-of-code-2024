package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []byte {
	return []byte(input.Lines()[0])
}

func process(diskMap []byte) uint64 {
	dm := unpack(diskMap)
	for l, r := 0, len(dm)-1; l < r; {
		if dm[l] >= 0 {
			l++
			continue
		}
		if dm[r] < 0 {
			r--
			continue
		}
		dm[l] = dm[r]
		dm[r] = -1
		l++
		r--
	}

	checksum := uint64(0)
	for i, id := range dm {
		if id < 0 {
			break
		}
		checksum += uint64(i * id)
	}

	return checksum
}

func unpack(diskMap []byte) []int {
	var unpackedDiskMap []int
	for i := 0; i < len(diskMap); i += 2 {
		blocks := diskMap[i] - '0'
		free := 0
		if i+1 < len(diskMap) {
			free = int(diskMap[i+1] - '0')
		}
		for range blocks {
			unpackedDiskMap = append(unpackedDiskMap, i/2)
		}
		for range free {
			unpackedDiskMap = append(unpackedDiskMap, -1)
		}
	}

	return unpackedDiskMap
}
