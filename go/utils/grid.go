package utils

import (
	"strconv"
	"strings"
)

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

type CharGrid struct {
	Grid   [][]string
	Width  int
	Height int
}

func CreateCharGridFromLines(lines []string) CharGrid {
	g := CharGrid{}
	for i, line := range lines {
		if i == 0 {
			g.Width = len(line)
		}
		if len(line) != g.Width {
			panic("non-rectangular grid")
		}
		row := SplitToChars(line)
		g.Grid = append(g.Grid, row)
		g.Height++
	}
	return g
}

func CreateEmptyCharGrid(width, height int) CharGrid {
	grid := CharGrid{
		Grid:   make([][]string, 0, height),
		Width:  width,
		Height: height}

	for i := 0; i < grid.Height; i++ {
		grid.Grid = append(grid.Grid, make([]string, grid.Width))
	}
	return grid
}

func (g CharGrid) SetValue(value string) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.Grid[y][x] = value
		}
	}
}

func (g *CharGrid) Find(x string) (row, col int) {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			if g.Grid[r][c] == x {
				return r, c
			}
		}
	}
	return -1, -1
}

func (g CharGrid) String() string {
	var rows []string
	for r := 0; r < g.Height; r++ {
		row := strings.Join(g.Grid[r], "")
		rows = append(rows, row)
	}
	return strings.Join(rows, "\n")
}

// InBounds - is (y, x) in grid
func (g CharGrid) InBounds(y, x int) bool {
	return 0 <= y && y < g.Height && 0 <= x && x < g.Width
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
			g.Grid[y][x] = value
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
