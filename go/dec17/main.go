package main

import (
	"fmt"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	winds := u.SplitToChars(lines[0])
	return stackTiles(winds, 2022)
}

func task2(lines []string) int {
	winds := u.SplitToChars(lines[0])
	period, periodHeight := findPeriod(winds)
	nrTiles := 1000000000000
	nrCycles := nrTiles / period
	remainder := nrTiles % period
	height := stackTiles(winds, remainder+period)
	height += (nrCycles - 1) * periodHeight
	return height
}

func findPeriod(winds []string) (period, height int) {
	heights := stackHeights(winds, 10000)
	o := 1000
	for p := 1; p < 4000; p++ {
		d := heights[o+p] - heights[o]
		i := 2
		period = p
		for {
			idx := o + i*p
			if idx >= len(heights) {
				break
			}
			md := heights[idx] - heights[o]
			if md != i*d {
				period = -1
				break
			}
			i++
		}
		if period != -1 {
			return period, d
		}
	}
	return -1, 0
}

func stackTiles(winds []string, n int) int {
	nrWinds := len(winds)
	pieces := makePieces()
	nrPieces := len(pieces)
	pieceNr := 0
	step := 0
	cv := newCave(600)
	for {
		p := pieces[pieceNr%nrPieces]
		r := cv.top + 4 // bottom row of new symbol
		c := 2
		cv.newCaveCeiling(r)
		for {
			wind := winds[step%nrWinds]
			step++
			switch wind {
			case ">":
				if canMoveHor(cv, p, r, c, +1) {
					c++
				}
			case "<":
				if canMoveHor(cv, p, r, c, -1) {
					c--
				}
			}
			if canMoveDown(cv, p, r, c) {
				r--
				continue
			}
			// Glue

			cv.addPiece(p, r, c)
			break
		}
		pieceNr++
		if pieceNr == n {
			//cv.PrintTop(10)
			break
		}
	}
	return cv.top + 1
}

func stackHeights(winds []string, n int) []int {
	nrWinds := len(winds)
	pieces := makePieces()
	nrPieces := len(pieces)
	pieceNr := 0
	step := 0
	cv := newCave(600)
	heights := make([]int, 0, n)
	for {
		p := pieces[pieceNr%nrPieces]
		r := cv.top + 4 // bottom row of new symbol
		c := 2
		cv.newCaveCeiling(r)
		for {
			wind := winds[step%nrWinds]
			step++
			switch wind {
			case ">":
				if canMoveHor(cv, p, r, c, +1) {
					c++
				}
			case "<":
				if canMoveHor(cv, p, r, c, -1) {
					c--
				}
			}
			if canMoveDown(cv, p, r, c) {
				r--
				continue
			}
			cv.addPiece(p, r, c)
			break
		}
		pieceNr++
		if pieceNr%1_000_000 == 0 {
			fmt.Printf("piece %d, wind=%d, top=%d\n", pieceNr, step, cv.top)
		}

		heights = append(heights, cv.top+1)

		if pieceNr == n {
			break
		}
	}
	return heights
}

type cave struct {
	width  int
	top    int
	bottom int
	g      [][]bool
}

func newCave(cap int) *cave {
	c := cave{
		width:  7,
		top:    -1,
		bottom: 0,
	}

	c.g = make([][]bool, 0, cap)
	for i := 0; i < cap; i++ {
		c.g = append(c.g, make([]bool, c.width))
	}
	return &c
}

func (c *cave) cap() int {
	return len(c.g)
}

func (c *cave) newCaveCeiling(r int) {
	margin := c.cap() + c.bottom - r
	if margin < 10 {
		c.addLines(10000)
	}
}

func (c *cave) addLines(n int) {
	for i := 0; i < n; i++ {
		c.g = append(c.g, make([]bool, c.width))
	}
}

func canMoveHor(c *cave, p piece, pr, pc, dir int) bool {
	pc += dir
	if pc < 0 {
		return false
	}
	if pc+p.w > c.width {
		return false
	}
	return !c.overlap(p, pr, pc)
}

func canMoveDown(c *cave, p piece, pr, pc int) bool {
	pr--
	if pr < 0 {
		return false
	}
	return !c.overlap(p, pr, pc)
}

func (cv *cave) overlap(p piece, pr, pc int) bool {
	for r := 0; r < p.h; r++ {
		caveRow := pr + r - cv.bottom
		for c := 0; c < p.w; c++ {
			caveCol := c + pc
			if caveRow < 0 {
				fmt.Println(r, c, pr, pc, cv.bottom, cv.top)
			}
			if p.g[r][c] && cv.g[caveRow][caveCol] {
				return true
			}
		}
	}
	return false
}

func (c *cave) PrintTop(nrLines int) {
	b := c.top - nrLines
	if c.bottom > b {
		b = c.bottom
	}
	for r := c.top; r >= b; r-- {
		cr := r - b
		m := "|"
		for col := 0; col < c.width; col++ {
			if c.g[cr][col] {
				m += "#"
			} else {
				m += "."
			}
		}
		m += "|"
		fmt.Println(m)
	}
	fmt.Println("+-------+")
}

// addPiece adds a piece and changes cv top and bottom if needed.
func (cv *cave) addPiece(p piece, pr, pc int) {
	bottomDiff := 0
	for r := 0; r < p.h; r++ {
		caveRow := r + pr - cv.bottom
		for c := 0; c < p.w; c++ {
			if p.g[r][c] {
				cv.g[caveRow][c+pc] = true
				if r+pr > cv.top {
					cv.top = r + pr
				}
			}
		}
		fullLine := true
		for x := 0; x < cv.width; x++ {
			if !cv.g[caveRow][x] {
				fullLine = false
				break
			}
		}
		if fullLine {
			bottomDiff = r + pr - cv.bottom
		}
	}
	if bottomDiff > 0 {
		cv.g = cv.g[bottomDiff:]
		cv.bottom += bottomDiff
	}
}

func makePieces() []piece {
	ps := make([]piece, 5)
	ps[0] = newPiece(4, 1)
	ps[0].setAll()

	ps[1] = newPiece(3, 3)
	ps[1].setAll()
	ps[1].g[0][0] = false
	ps[1].g[0][2] = false
	ps[1].g[2][0] = false
	ps[1].g[2][2] = false

	ps[2] = newPiece(3, 3)
	ps[2].setAll()
	for r := 1; r < 3; r++ {
		for c := 0; c < 2; c++ {
			ps[2].g[r][c] = false
		}
	}

	ps[3] = newPiece(1, 4)
	ps[3].setAll()

	ps[4] = newPiece(2, 2)
	ps[4].setAll()
	return ps
}

type piece struct {
	g [][]bool
	w int
	h int
}

func (p *piece) setAll() {
	for r := 0; r < p.h; r++ {
		for c := 0; c < p.w; c++ {
			p.g[r][c] = true
		}
	}
}

func newPiece(width, height int) piece {
	g := piece{
		g: make([][]bool, 0, height),
		w: width,
		h: height}

	for i := 0; i < g.h; i++ {
		g.g = append(g.g, make([]bool, g.w))
	}
	return g
}
