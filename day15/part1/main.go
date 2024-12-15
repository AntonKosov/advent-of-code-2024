package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2024/aoc/input"
	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (warehouse Warehouse, movements []math.Vector2[int]) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	warehouse.boxes = map[math.Vector2[int]]bool{}
	for row := 0; lines[0] != ""; row++ {
		line := lines[0]
		lines = lines[1:]
		warehouse.walls = append(warehouse.walls, make([]bool, len(line)))
		for col, v := range line {
			switch v {
			case '#':
				warehouse.walls[row][col] = true
			case '@':
				warehouse.robotPosition = math.NewVector2(col, row)
			case 'O':
				warehouse.boxes[math.NewVector2(col, row)] = true
			}
		}
	}

	dirs := map[rune]math.Vector2[int]{
		'^': math.NewVector2(0, -1),
		'v': math.NewVector2(0, 1),
		'>': math.NewVector2(1, 0),
		'<': math.NewVector2(-1, 0),
	}
	for _, line := range lines {
		for _, d := range line {
			movements = append(movements, dirs[d])
		}
	}

	return warehouse, movements
}

func process(warehouse Warehouse, movements []math.Vector2[int]) int {
	for _, dir := range movements {
		warehouse.move(dir)
	}

	return warehouse.GPSSum()
}

type Warehouse struct {
	robotPosition math.Vector2[int]
	walls         [][]bool
	boxes         map[math.Vector2[int]]bool
}

func (w *Warehouse) move(dir math.Vector2[int]) {
	pos := w.robotPosition.Add(dir)
	if w.walls[pos.Y][pos.X] {
		return
	}

	if w.boxes[pos] && !w.moveBox(pos, dir) {
		return
	}

	w.robotPosition = pos
}

func (w *Warehouse) moveBox(pos, dir math.Vector2[int]) bool {
	box := w.boxes[pos]
	nextPos := pos.Add(dir)
	if w.walls[nextPos.Y][nextPos.X] {
		return false
	}

	if w.boxes[nextPos] && !w.moveBox(nextPos, dir) {
		return false
	}

	delete(w.boxes, pos)
	w.boxes[nextPos] = box

	return true
}

func (w *Warehouse) GPSSum() int {
	sum := 0
	for pos := range w.boxes {
		sum += pos.Y*100 + pos.X
	}

	return sum
}
