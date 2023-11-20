package life

import (
	"math/rand"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	cells := make([][]bool, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) RandInit(percents int) {
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			w.Cells[i][j] = rand.Int()%100 <= percents
		}
	}
}

func NextState(currentWorld, nextWorld *World) *World {
	for y := 0; y < currentWorld.Height; y++ {
		for x := 0; x < currentWorld.Width; x++ {
			cell := currentWorld.Cells[y][x]
			n := currentWorld.neighbors(x, y)
			if !cell && n == 3 {
				nextWorld.Cells[y][x] = true
				continue
			}
			if cell && n != 2 && n != 3 {
				nextWorld.Cells[y][x] = false
				continue
			}
			nextWorld.Cells[y][x] = cell
		}
	}
	return nextWorld
}

func (w World) neighbors(x, y int) int {
	var res int
	st_y, fn_y := max(y-1, 0), min(w.Height-1, y+1)
	st_x, fn_x := max(x-1, 0), min(w.Width-1, x+1)
	for i := st_y; i <= fn_y; i++ {
		for j := st_x; j <= fn_x; j++ {
			if i == y && j == x {
				continue
			}
			if w.Cells[i][j] {
				res++
			}
		}
	}
	return res
}
