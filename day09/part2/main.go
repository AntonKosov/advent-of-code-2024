package main

import (
	"container/list"
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
	fragments := transformToList(diskMap)
	defragment(fragments)

	return calcChecksum(fragments)
}

func calcChecksum(fragments *list.List) uint64 {
	checksum := uint64(0)
	for pos, e := 0, fragments.Front(); e != nil; e = e.Next() {
		fr := e.Value.(*Fragment)
		if fr.id < 0 {
			pos += fr.freeBlocks
			continue
		}
		for range fr.blocks {
			checksum += uint64(fr.id * pos)
			pos++
		}
		pos += fr.freeBlocks
	}

	return checksum
}

func defragment(fragments *list.List) {
	for back := fragments.Back(); back != nil; back = back.Prev() {
		backEl := back.Value.(*Fragment)
		for front := fragments.Front(); front != nil && front != back; front = front.Next() {
			frontEl := front.Value.(*Fragment)
			if fb := frontEl.freeBlocks; fb >= backEl.blocks {
				newFr := Fragment{
					id:         backEl.id,
					blocks:     backEl.blocks,
					freeBlocks: fb - backEl.blocks,
				}
				frontEl.freeBlocks = 0
				backEl.id = -1
				backEl.freeBlocks += backEl.blocks
				backEl.blocks = 0
				fragments.InsertAfter(&newFr, front)
				break
			}
		}
	}
}

func transformToList(diskMap []byte) *list.List {
	fragments := transformToFragments(diskMap)
	fragmentList := list.New()
	for _, f := range fragments {
		fragmentList.PushBack(&f)
	}

	return fragmentList
}

func transformToFragments(diskMap []byte) []Fragment {
	unpackedDiskMap := make([]Fragment, 0, 1+len(diskMap)+1)
	for i := 0; i < len(diskMap); i += 2 {
		blocks := diskMap[i] - '0'
		free := 0
		if i+1 < len(diskMap) {
			free = int(diskMap[i+1] - '0')
		}
		unpackedDiskMap = append(unpackedDiskMap, Fragment{
			id:         i / 2,
			blocks:     int(blocks),
			freeBlocks: free,
		})
	}

	return unpackedDiskMap
}

type Fragment struct {
	id         int
	blocks     int
	freeBlocks int
}
