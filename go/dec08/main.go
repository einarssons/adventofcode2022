package main

import (
	"fmt"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	grid := u.CreateDigitGridFromLines(lines)
	fmt.Printf("width: %d, height: %d\n", grid.Width, grid.Height)
	nrVisible := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			v := visible(x, y, &grid)
			//fmt.Printf("(%d, %d, %d\n", y, x, v)
			nrVisible += v
		}

	}
	return nrVisible
}

func task2(lines []string) int {
	grid := u.CreateDigitGridFromLines(lines)
	maxView := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			totView := view(x, y, &grid)
			if totView > maxView {
				maxView = totView
			}
		}
	}
	return maxView
}

func visible(col, row int, grid *u.DigitGrid) int {
	if col == 0 || row == 0 || col == grid.Width-1 || row == grid.Height-1 {
		return 1
	}
	height := grid.Grid[row][col]
	// Same row from left
	maxHeight := 0
	for x := 0; x < col; x++ {
		if grid.Grid[row][x] > maxHeight {
			maxHeight = grid.Grid[row][x]
		}
	}
	if maxHeight < height {
		return 1
	}
	// Same row from right
	maxHeight = 0
	for x := col + 1; x < grid.Width; x++ {
		if grid.Grid[row][x] > maxHeight {
			maxHeight = grid.Grid[row][x]
		}
	}
	if maxHeight < height {
		return 1
	}
	// Same col from top
	maxHeight = 0
	for y := 0; y < row; y++ {
		if grid.Grid[y][col] > maxHeight {
			maxHeight = grid.Grid[y][col]
		}
	}
	if maxHeight < height {
		return 1
	}
	// Same col from bottom
	maxHeight = 0
	for y := row + 1; y < grid.Height; y++ {
		if grid.Grid[y][col] > maxHeight {
			maxHeight = grid.Grid[y][col]
		}
	}
	if maxHeight < height {
		return 1
	}
	return 0
}

func view(col, row int, grid *u.DigitGrid) int {
	totalView := 1
	// Same row to left
	height := grid.Grid[row][col]
	nrTrees := 0
	for x := col - 1; x >= 0; x-- {
		if grid.Grid[row][x] < height {
			nrTrees++
			continue
		}
		if grid.Grid[row][x] >= height {
			nrTrees++
			break
		}
	}
	totalView *= nrTrees
	// Same row to right
	nrTrees = 0
	for x := col + 1; x < grid.Width; x++ {
		if grid.Grid[row][x] < height {
			nrTrees++
			continue
		}
		if grid.Grid[row][x] >= height {
			nrTrees++
			break
		}
	}
	totalView *= nrTrees
	// Same col to top
	nrTrees = 0
	for y := row - 1; y >= 0; y-- {
		if grid.Grid[y][col] < height {
			nrTrees++
			continue
		}
		if grid.Grid[y][col] >= height {
			nrTrees++
			break
		}
	}
	totalView *= nrTrees
	// Same col to bottom
	nrTrees = 0
	for y := row + 1; y < grid.Height; y++ {
		if grid.Grid[y][col] < height {
			nrTrees++
			continue
		}
		if grid.Grid[y][col] >= height {
			nrTrees++
			break
		}
	}
	totalView *= nrTrees
	return totalView
}
