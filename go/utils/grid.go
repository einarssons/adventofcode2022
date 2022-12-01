package utils

import "strconv"

type DigitGrid struct {
	Grid   [][]int
	Width  int
	Height int
}

func CreateDigitGridFromLines(lines []string) DigitGrid {
	g := DigitGrid{}
	for i, line := range lines {
		if i == 0 {
			g.Width = len(line)
		}
		if len(line) != g.Width {
			panic("non-rectangular grid")
		}
		row := make([]int, 0, g.Width)
		digits := SplitToChars(line)
		for _, digit := range digits {
			nr, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}
			row = append(row, nr)
		}
		g.Grid = append(g.Grid, row)
		g.Height++
	}
	return g
}

func CreateZeroDigitGrid(width, height int) DigitGrid {
	grid := DigitGrid{
		Grid:   make([][]int, 0, height),
		Width:  width,
		Height: height}

	for i := 0; i < grid.Height; i++ {
		grid.Grid = append(grid.Grid, make([]int, grid.Width))
	}
	return grid
}

// InBounds - is (y, x) in grid
func (g DigitGrid) InBounds(y, x int) bool {
	return 0 <= y && y < g.Height && 0 <= x && x < g.Width
}

func (g DigitGrid) SetValue(value int) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.Grid[x][y] = value
		}
	}
}

type BoolGrid struct {
	Grid   [][]bool
	Width  int
	Height int
}

func CreateBoolGrid(width, height int) BoolGrid {
	grid := BoolGrid{
		Grid:   make([][]bool, 0, height),
		Width:  width,
		Height: height}

	for i := 0; i < grid.Height; i++ {
		grid.Grid = append(grid.Grid, make([]bool, grid.Width))
	}
	return grid
}
