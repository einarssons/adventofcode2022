package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadRawLinesFromFile("input")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines, 50, false))
}

func task1(lines []string) int {
	grid := makeGridFromLines(lines[:len(lines)-2], 4)
	dirs := findDirs(lines[len(lines)-1])
	facing := 0
	p := grid.findFirst()
	for _, d := range dirs {
		switch d.turn {
		case "L":
			facing = (facing - 1 + 4) % 4
		case "R":
			facing = (facing + 1) % 4
		default:
			switch facing {
			case 0:
				p = grid.hor(p, d.nr, 1)
			case 1:
				p = grid.ver(p, d.nr, 1)
			case 2:
				p = grid.hor(p, d.nr, -1)
			default: // 3
				p = grid.ver(p, d.nr, -1)
			}
			//fmt.Printf("(%d, %d)\n", p.r, p.c)
		}
	}
	points := 1000*(p.r+1) + 4*(p.c+1) + facing
	return points
}

func task2(lines []string, side int, test bool) int {
	grid := makeGridFromLines(lines[:len(lines)-2], side)
	dirs := findDirs(lines[len(lines)-1])
	facing := 0 // facing
	p := grid.findFirst()
	for _, d := range dirs {
		switch d.turn {
		case "L":
			facing = (facing - 1 + 4) % 4
		case "R":
			facing = (facing + 1) % 4
		default:
			p, facing = grid.moveOnCube(p, d.nr, facing, test)
		}
	}
	points := 1000*(p.r+1) + 4*(p.c+1) + facing
	return points
}

type grid struct {
	g    []string
	side int
	w    int
	h    int
}

type pos struct {
	r, c int
}

func makeGridFromLines(lines []string, side int) *grid {
	w := 0
	for _, l := range lines {
		if len(l) > w {
			w = len(l)
		}
	}
	return &grid{g: lines, side: side, w: w, h: len(lines)}
}

func (g *grid) findFirst() pos {
	row := g.g[0]
	for c := 0; c < len(row); c++ {
		char := string(row[c])
		if char == "." {
			return pos{0, c}
		}
	}
	panic("no start")
}

func (g *grid) nextHor(p pos, dir int) pos {
	row := g.g[p.r]
	lRow := len(row)
	c := p.c
	for {
		n := (c + dir + lRow) % lRow
		nextChar := string(row[n])
		if nextChar == " " {
			c = n
			continue
		}
		return pos{p.r, n}
	}
}

func (g *grid) hor(p pos, steps, dir int) pos {
	row := g.g[p.r]
	for i := 0; i < steps; i++ {
		np := g.nextHor(p, dir)
		nextChar := string(row[np.c])
		if nextChar == "#" {
			return pos{p.r, p.c}
		}
		p = np
	}
	return p
}

func (g *grid) char(p pos) string {
	return string(g.g[p.r][p.c])
}

func (g *grid) moveOnCube(p pos, steps, facing int, test bool) (pos, int) {
	next := g.nextInput
	if test {
		next = g.nextTest
	}
	for i := 0; i < steps; i++ {
		np, nf := next(p, facing)
		nextChar := g.char(np)
		if nextChar == "#" {
			return p, facing
		}
		p, facing = np, nf
	}
	return p, facing
}

func (g *grid) in(p pos) bool {
	if p.r < 0 || p.c < 0 {
		return false
	}
	if p.r >= g.h {
		return false
	}
	if p.c < len(g.g[p.r]) && g.char(p) != " " {
		return true
	}
	return false
}

func (g *grid) nextInput(p pos, f int) (pos, int) {
	switch f {
	case 0:
		if g.in(pos{p.r, p.c + 1}) {
			return pos{p.r, p.c + 1}, f
		}
	case 1:
		if g.in(pos{p.r + 1, p.c}) {
			return pos{p.r + 1, p.c}, f
		}
	case 2:
		if g.in(pos{p.r, p.c - 1}) {
			return pos{p.r, p.c - 1}, f
		}
	case 3:
		if g.in(pos{p.r - 1, p.c}) {
			return pos{p.r - 1, p.c}, f
		}
	}
	s := g.side
	switch f {
	case 0: // right
		rs := p.r / g.side
		relRow := p.r - rs*s
		switch rs {
		case 0:
			return pos{3*s - 1 - relRow, 2*s - 1}, 2 // 2 -> 5 left
		case 1:
			return pos{s - 1, 2*s + relRow}, 3 // 3-> 2 up
		case 2:
			return pos{s - 1 - relRow, 3*s - 1}, 2 //5 -> 2 left
		case 3:
			return pos{3*s - 1, s + relRow}, 3 // 6 -> 5 up
		default:
			panic("dir 0")
		}
	case 1: // down
		rc := p.c / g.side
		relCol := p.c - rc*g.side
		switch rc {
		case 0:
			return pos{0, 2*s + relCol}, 1 // 6 -> 2 down
		case 1:
			return pos{3*s + relCol, s - 1}, 2 // 5 -> 6 left
		case 2:
			return pos{s + relCol, 2*s - 1}, 2 // 2 -> 3 left
		default:
			panic("dir 1")
		}
	case 2: // left
		rs := p.r / g.side
		relRow := p.r - rs*g.side
		switch rs {
		case 0:
			return pos{3*s - 1 - relRow, 0}, 0 // 1 -> 4 right
		case 1:
			return pos{2 * s, relRow}, 1 // 3-> 4 down
		case 2:
			return pos{s - 1 - relRow, s}, 0 // 4-> 1 right
		case 3:
			return pos{0, s + relRow}, 1 // 6 -> 1 down
		default:
			panic("dir 2")
		}
	default: // 3 up
		rc := p.c / g.side
		relCol := p.c - rc*g.side
		switch rc {
		case 0:
			return pos{s + relCol, s}, 0 // 4-> 3 right
		case 1:
			return pos{3*s + relCol, 0}, 0 // 1 -> 6 right
		case 2:
			return pos{g.h - 1, relCol}, 3 // 2-> 6 up
		default:
			panic("dir 3")
		}
	}
}

func (g *grid) nextTest(p pos, f int) (pos, int) {
	switch f {
	case 0:
		if g.in(pos{p.r, p.c + 1}) {
			return pos{p.r, p.c + 1}, f
		}
	case 1:
		if g.in(pos{p.r + 1, p.c}) {
			return pos{p.r + 1, p.c}, f
		}
	case 2:
		if g.in(pos{p.r, p.c - 1}) {
			return pos{p.r, p.c - 1}, f
		}
	case 3:
		if g.in(pos{p.r - 1, p.c}) {
			return pos{p.r - 1, p.c}, f
		}
	}
	switch f {
	case 0: // right
		rs := p.r / g.side
		relRow := p.r - rs*g.side
		switch rs {
		case 0:
			return pos{g.h - 1 - relRow, g.w - 1}, 2 // 1 -> 6 left
		case 1:
			return pos{2 * g.side, g.w - 1 - relRow}, 1 // 4-> 6 down
		case 2:
			return pos{g.side - 1 - relRow, 3*g.side - 1}, 2 // 6-> 1 left
		default:
			panic("dir 0")
		}
	case 1: // down
		rc := p.c / g.side
		relCol := p.c - rc*g.side
		switch rc {
		case 0:
			return pos{g.h - 1, 3*g.side - 1 - relCol}, 3 // 2 -> 5 up
		case 1:
			return pos{g.h - 1 - relCol, 2 * g.side}, 0 // 3 -> 5 right
		case 2:
			return pos{g.side - 1 - relCol, 2*g.side - 1}, 3 //5 down -> 2 up
		case 3:
			return pos{2*g.side - 1 - relCol, 0}, 0 // 6 -> 2 right
		default:
			panic("dir 1")
		}
	case 2: // left
		rs := p.r / g.side
		relRow := p.r - rs*g.side
		switch rs {
		case 0:
			return pos{g.side, g.side + relRow}, 1 // 1 -> 3 down
		case 1:
			return pos{g.h, g.w - relRow}, 3 // 2-> 6 up
		case 2:
			return pos{2*g.side - 1, 2*g.side - 1 - relRow}, 3 // 5-> 3 up
		default:
			panic("dir 2")
		}
	default: // 3 up
		rc := p.c / g.side
		relCol := p.c - rc*g.side
		switch rc {
		case 0:
			return pos{0, 3*g.side - 1 - relCol}, 1 // 2-> 1 down
		case 1:
			return pos{relCol, 2 * g.side}, 0 // 3 -> 1 right
		case 2:
			return pos{g.side - 1 - relCol, g.side}, 1 // 1-> 2 down
		case 3:
			return pos{2*g.side - 1 - relCol, 3*g.side - 1}, 2 // 6 -> 4 left
		default:
			panic("dir 3")
		}
	}
}

func (g *grid) nextVer(p pos, dir int) pos {
	h := len(g.g)
	r := p.r
	for {
		n := (r + dir + h) % h
		// fmt.Println(n, p.c, len(g.g[n]))
		if len(g.g[n]) < p.c+1 {
			r = n
			continue
		}
		nextChar := string(g.g[n][p.c])
		if nextChar == " " {
			r = n
			continue
		}
		return pos{n, p.c}
	}
}

func (g *grid) ver(p pos, steps, dir int) pos {
	for i := 0; i < steps; i++ {
		np := g.nextVer(p, dir)
		nextChar := string(g.g[np.r][np.c])
		if nextChar == "#" {
			return pos{p.r, p.c}
		}
		p = np
	}
	return p
}

type dir struct {
	nr   int
	turn string
}

func findDirs(line string) []dir {
	var dirs []dir
	p := 0
	for {
		lPos := strings.Index(line[p:], "L")
		rPos := strings.Index(line[p:], "R")
		switch {
		case lPos == -1 && rPos == -1:
			if p < len(line) {
				dirs = append(dirs, dir{nr: u.Atoi(line[p:])})
				p = len(line)
			}
		case lPos == -1 || rPos == -1:
			dPos := lPos
			if dPos == -1 {
				dPos = rPos
			}
			if dPos == 0 {
				dirs = append(dirs, dir{turn: string(line[p])})
				p++
			} else {
				nr := u.Atoi(line[p : p+dPos])
				dirs = append(dirs, dir{nr: nr})
				p += dPos
			}
		default: // both found
			mPos := lPos
			if rPos < mPos {
				mPos = rPos
			}
			if mPos == 0 {
				dirs = append(dirs, dir{turn: string(line[p])})
				p++
			} else {
				nr := u.Atoi(line[p : p+mPos])
				dirs = append(dirs, dir{nr: nr})
				p += mPos
			}
		}
		if p == len(line) {
			break
		}
	}
	return dirs
}
