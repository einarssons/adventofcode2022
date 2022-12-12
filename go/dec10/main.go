package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines))
	s := task2(lines)
	fmt.Println("task2:")
	fmt.Println(s)
}

var thresholds = []int{20, 60, 100, 140, 180, 220}

func probeValue(value, cycle int) int {
	if u.ContainsInt(cycle, thresholds) {
		return value * cycle
	}
	return 0
}

func task1(lines []string) int {
	sum := 0
	cycle := 0
	X := 1
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "noop"):
			cycle++
			sum += probeValue(X, cycle)
		case strings.HasPrefix(line, "addx"):
			_, val, _ := strings.Cut(line, " ")
			steps := u.Atoi(val)
			cycle++
			sum += probeValue(X, cycle)
			cycle++
			sum += probeValue(X, cycle)
			X += steps
		}
	}
	return sum
}

type screen struct {
	nrRows int
	nrCols int
	pixels []string
}

func createScreen(rows, cols int) *screen {
	s := screen{nrRows: rows, nrCols: cols}
	s.pixels = make([]string, s.nrRows*s.nrCols)
	for i := 0; i < s.nrRows*s.nrCols; i++ {
		s.pixels[i] = "_"
	}
	return &s
}

func (s *screen) draw(x, cycle int) {
	// sprite [x-1, x, x+1]
	pos := x
	pixel := (cycle - 1)
	col := pixel % s.nrCols
	if pos-1 <= col && col <= pos+1 {
		s.pixels[pixel] = "#"
	} else {
		s.pixels[pixel] = " "
	}
}

func (s *screen) String() string {
	rows := make([]string, s.nrRows)
	for row := 0; row < s.nrRows; row++ {
		line := strings.Join(s.pixels[row*s.nrCols:(row+1)*s.nrCols], "")
		rows[row] = line
	}
	return strings.Join(rows, "\n")
}

func task2(lines []string) string {
	s := createScreen(6, 40)
	cycle := 0
	X := 1
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "noop"):
			cycle++
			s.draw(X, cycle)
		case strings.HasPrefix(line, "addx"):
			_, val, _ := strings.Cut(line, " ")
			steps := u.Atoi(val)
			cycle++
			s.draw(X, cycle)
			cycle++
			s.draw(X, cycle)
			X += steps
		}
	}
	return s.String()
}
