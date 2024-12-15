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

var (
	DirLeft  = math.NewVector2(-1, 0)
	DirRight = math.NewVector2(1, 0)
	DirUp    = math.NewVector2(0, -1)
	DirDown  = math.NewVector2(0, 1)
)

func read() (warehouse Warehouse, movements []math.Vector2[int]) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	warehouse.boxes = map[math.Vector2[int]]Box{}
	for row := 0; lines[0] != ""; row++ {
		line := lines[0]
		lines = lines[1:]
		warehouse.walls = append(warehouse.walls, make([]bool, 2*len(line)))
		for col, v := range line {
			col *= 2
			switch v {
			case '#':
				warehouse.walls[row][col] = true
				warehouse.walls[row][col+1] = true
			case '@':
				warehouse.robotPosition = math.NewVector2(col, row)
			case 'O':
				warehouse.boxes[math.NewVector2(col, row)] = Box{connection: DirRight}
				warehouse.boxes[math.NewVector2(col+1, row)] = Box{connection: DirLeft}
			}
		}
	}

	dirs := map[rune]math.Vector2[int]{'^': DirUp, 'v': DirDown, '>': DirRight, '<': DirLeft}
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

type Box struct {
	connection math.Vector2[int]
}

type BoxMovement struct {
	from, to math.Vector2[int]
}

type Warehouse struct {
	robotPosition math.Vector2[int]
	walls         [][]bool
	boxes         map[math.Vector2[int]]Box
}

func (w *Warehouse) wall(pos math.Vector2[int]) bool { return w.walls[pos.Y][pos.X] }

func (w *Warehouse) move(dir math.Vector2[int]) {
	pos := w.robotPosition.Add(dir)
	if w.wall(pos) {
		return
	}

	if _, ok := w.boxes[pos]; ok && !w.moveBox(pos, dir) {
		return
	}

	w.robotPosition = pos
}

func (w *Warehouse) moveBox(pos, dir math.Vector2[int]) bool {
	var movements []BoxMovement
	if !w.prepareMovements(pos, dir, map[math.Vector2[int]]bool{}, &movements) {
		return false
	}

	for _, movement := range movements {
		fromPos := movement.from
		box := w.boxes[fromPos]
		delete(w.boxes, fromPos)
		w.boxes[movement.to] = box
	}

	return true
}

func (w *Warehouse) prepareMovements(
	pos, dir math.Vector2[int], movedBoxes map[math.Vector2[int]]bool, movements *[]BoxMovement,
) (canBeMoved bool) {
	if movedBoxes[pos] {
		return true
	}

	box := w.boxes[pos]

	leftPos, rightPos := pos, pos.Add(box.connection)
	if leftPos.X > rightPos.X {
		leftPos, rightPos = rightPos, leftPos
	}

	defer func() {
		if canBeMoved {
			movedBoxes[leftPos] = true
			movedBoxes[rightPos] = true
		}
	}()

	newLeftPos := leftPos.Add(dir)
	newRightPos := rightPos.Add(dir)

	if w.wall(newLeftPos) || w.wall(newRightPos) {
		return false
	}

	move := func(to math.Vector2[int], boxMovements ...BoxMovement) bool {
		if _, ok := w.boxes[to]; ok && !w.prepareMovements(to, dir, movedBoxes, movements) {
			return false
		}

		*movements = append(*movements, boxMovements...)

		return true
	}

	switch dir {
	case DirRight:
		return move(newRightPos, BoxMovement{from: rightPos, to: newRightPos}, BoxMovement{from: leftPos, to: newLeftPos})
	case DirLeft:
		return move(newLeftPos, BoxMovement{from: leftPos, to: newLeftPos}, BoxMovement{from: rightPos, to: newRightPos})
	}

	return move(newLeftPos, BoxMovement{from: leftPos, to: newLeftPos}) &&
		move(newRightPos, BoxMovement{from: rightPos, to: newRightPos})
}

func (w *Warehouse) GPSSum() int {
	sum := 0
	for pos, box := range w.boxes {
		if box.connection == DirRight {
			sum += pos.Y*100 + pos.X
		}
	}

	return sum
}

// func (w *Warehouse) print() {
// 	for r, row := range w.walls {
// 		for c, wall := range row {
// 			ch := ' '
// 			pos := math.NewVector2(c, r)
// 			if wall {
// 				ch = '#'
// 			} else if w.robotPosition == pos {
// 				ch = '@'
// 			} else if box := w.boxes[pos]; box != nil {
// 				ch = '['
// 				if box.connection.X < 0 {
// 					ch = ']'
// 				}
// 			}
// 			fmt.Print(string(ch))
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }
