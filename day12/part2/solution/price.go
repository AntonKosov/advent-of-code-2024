package solution

import (
	"slices"

	"github.com/AntonKosov/advent-of-code-2024/aoc/math"
	"github.com/AntonKosov/advent-of-code-2024/aoc/slice"
)

func Parse(lines []string) [][]int {
	return slice.Map(lines, func(line string) []int {
		return slice.Map([]rune(line), func(r rune) int { return int(r) })
	})
}

func Price(garden [][]int) int {
	garden = addBorder(garden)
	id := 0
	price := 0
	for y, row := range garden {
		for x := range row {
			id--
			area, sides := measure(garden, x, y, id)
			price += area * sides
		}
	}

	return price
}

func addBorder(garden [][]int) [][]int {
	width, height := len(garden[0])+2, len(garden)+2
	gardenWithBorder := make([][]int, height)
	for y, row := range garden {
		br := make([]int, width)
		copy(br[1:], row)
		gardenWithBorder[y+1] = br
	}

	gardenWithBorder[0] = make([]int, width)
	gardenWithBorder[height-1] = make([]int, width)

	return gardenWithBorder
}

func measure(garden [][]int, x, y, id int) (area, sides int) {
	plotType := garden[y][x]
	if plotType <= 0 {
		return 0, 0
	}

	var edges Edges
	positions := []math.Vector2[int]{math.NewVector2(x, y)}
	for len(positions) > 0 {
		pos := positions[len(positions)-1]
		positions = positions[:len(positions)-1]
		if garden[pos.Y][pos.X] <= 0 {
			continue
		}

		garden[pos.Y][pos.X] = id
		area++
		check := func(dx, dy int, edge *[]math.Vector2[int]) {
			nextPos := pos.Add(math.NewVector2(dx, dy))
			plot := garden[nextPos.Y][nextPos.X]
			if plot != plotType && plot != id {
				*edge = append(*edge, pos)
				return
			}
			positions = append(positions, nextPos)
		}
		check(-1, 0, &edges.left)
		check(1, 0, &edges.right)
		check(0, -1, &edges.top)
		check(0, 1, &edges.bottom)
	}

	return area, edges.Count()
}

type Edges struct {
	left, right, top, bottom []math.Vector2[int]
}

func (e Edges) Count() int {
	return verticalSides(e.left) +
		verticalSides(e.right) +
		horizontalSides(e.top) +
		horizontalSides(e.bottom)
}

func verticalSides(positions []math.Vector2[int]) int {
	slices.SortFunc(positions, func(p1, p2 math.Vector2[int]) int {
		if dx := p1.X - p2.X; dx != 0 {
			return dx
		}

		return p1.Y - p2.Y
	})

	sides := 0
	for len(positions) > 0 {
		pos := positions[0]
		count := len(positions)
		for i := 1; i < count; i++ {
			if positions[i] != math.NewVector2(pos.X, pos.Y+i) {
				count = i
				break
			}
		}
		sides++
		positions = positions[count:]
	}

	return sides
}

func horizontalSides(positions []math.Vector2[int]) int {
	slices.SortFunc(positions, func(p1, p2 math.Vector2[int]) int {
		if dy := p1.Y - p2.Y; dy != 0 {
			return dy
		}

		return p1.X - p2.X
	})

	sides := 0
	for len(positions) > 0 {
		pos := positions[0]
		count := len(positions)
		for i := 1; i < count; i++ {
			if positions[i] != math.NewVector2(pos.X+i, pos.Y) {
				count = i
				break
			}
		}
		sides++
		positions = positions[count:]
	}

	return sides
}
