package main

import (
	"fmt"
	"sort"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	lines = u.TrimTrailingNewline(lines)
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

var shortest int
var cycle int
var prevMoves []map[pos]int
var start pos
var target pos

func task1(lines []string) int {
	shortest := end2end(lines, 0, true)
	return shortest
}

func task2(lines []string) int {
	total := 0
	shortest := end2end(lines, 0, true)
	fmt.Printf("Down steps: %d\n", shortest)
	total += shortest

	shortest = end2end(lines, total, false)
	fmt.Printf("Up steps: %d\n", shortest)
	total += shortest

	shortest = end2end(lines, total, true)
	fmt.Printf("Down again steps: %d\n", shortest)
	total += shortest

	return total
}

func end2end(lines []string, startIter int, forward bool) int {
	initShortest := 400
	for {
		v := parse(lines)
		if forward {
			start = pos{-1, 0}
			target = pos{v.h, v.w - 1}
		} else {
			start = pos{v.h, v.w - 1}
			target = pos{-1, 0}
		}
		shortest = initShortest
		cycle = v.h * v.w / u.GCD(v.h, v.w)
		initPrevMoves(cycle) // Remember previous modes mod cycle, and avoid them
		prevMoves[0][start] = 0
		for i := 0; i < startIter; i++ {
			v.stepBlizzards()
		}
		//v.print(start)
		v.next(start, 0)
		if shortest < initShortest {
			fmt.Printf("really shortest: %d\n", shortest)
			break
		}
		initShortest += 100
		fmt.Printf("increased shortest to %d\n", initShortest)

	}
	return shortest
}

type valley struct {
	h, w int
	bzs  []blizzard
}

func (v *valley) next(p pos, nr int) {
	nr++
	if nr >= shortest {
		return
	}
	v.stepBlizzards()
	//v.print(pos{-1, -1})
	dirs := v.freeDirs(p, nr)
	//fmt.Printf("nr %d, (%d, %d), %v\n", nr, p.r, p.c, dirs)

	for _, d := range dirs {
		n := p.move(d)
		addPrevMove(n, nr)
		if n == target {
			if nr < shortest {
				shortest = nr
				fmt.Printf("shortest %d\n", shortest)
			}
			break
		}
		v.next(n, nr)
	}
	v.backBlizzards()
}

func (v valley) inBounds(p pos, d dir) bool {
	if p == target {
		return true
	}
	zeroDir := dir{0, 0}
	if p == start && d == zeroDir {
		return true
	}
	return p.r >= 0 && p.r < v.h && p.c >= 0 && p.c < v.w
}

func (v valley) print(mp pos) {
	i := 0
	for r := 0; r < v.h; r++ {
		for c := 0; c < v.w; c++ {
			p := pos{r, c}
			if mp == p {
				fmt.Printf("E")
				continue
			}
			count := 0
			for j := i; j < len(v.bzs); j++ {
				if v.bzs[j].p == p {
					count++
				}
			}
			switch count {
			case 0:
				fmt.Printf(".")
			case 1:
				fmt.Printf(v.bzs[i].dir)
			default:
				fmt.Printf("%d", count)
			}
			i += count
		}
		fmt.Println()
	}
	fmt.Println()
}

func initPrevMoves(cycle int) {
	prevMoves = make([]map[pos]int, cycle)
	for i := 0; i < cycle; i++ {
		prevMoves[i] = make(map[pos]int)
	}
}

func isPrevMove(p pos, nr int) bool {
	cycleNr := nr % cycle
	pm := prevMoves[cycleNr]
	pNr, ok := pm[p]
	if ok && pNr <= nr {
		return true
	}
	return false
}

func addPrevMove(p pos, nr int) {
	cycleNr := nr % cycle
	prevMoves[cycleNr][p] = nr
}

func (v valley) freeDirs(p pos, nr int) []dir {
	dirs := make([]dir, 0, 5)
	for i := 0; i < 5; i++ {
		d := dir{0, 0}
		switch i {
		case 0: // Right
			d.c += 1
		case 1: // Down
			d.r += 1
		case 2: // Left
			d.c--
		case 3: // Up
			d.r--
		case 4: // Stay
		}
		n := p.move(d)
		if !v.inBounds(n, d) {
			continue
		}
		if isPrevMove(n, nr) {
			continue
		}
		if !v.hasBlizzard(n) {
			dirs = append(dirs, d)
		}
	}
	return dirs
}

func (v valley) hasBlizzard(p pos) bool {
	idx := sort.Search(len(v.bzs), func(i int) bool {
		rowDiff := v.bzs[i].p.r - p.r
		if rowDiff > 0 {
			return true
		}
		if rowDiff == 0 && v.bzs[i].p.c >= p.c {
			return true
		}
		return false
	})
	if idx == len(v.bzs) {
		return false
	}
	return v.bzs[idx].p == p
}

func (v *valley) stepBlizzards() {
	for i := range v.bzs {
		switch v.bzs[i].dir {
		case ">":
			v.bzs[i].p.c = (v.bzs[i].p.c + 1) % v.w
		case "<":
			v.bzs[i].p.c = (v.bzs[i].p.c - 1 + v.w) % v.w
		case "v":
			v.bzs[i].p.r = (v.bzs[i].p.r + 1) % v.h
		case "^":
			v.bzs[i].p.r = (v.bzs[i].p.r - 1 + v.h) % v.h
		}
	}
	v.sort()
}

func (v *valley) backBlizzards() {
	for i := range v.bzs {
		switch v.bzs[i].dir {
		case "<":
			v.bzs[i].p.c = (v.bzs[i].p.c + 1) % v.w
		case ">":
			v.bzs[i].p.c = (v.bzs[i].p.c - 1 + v.w) % v.w
		case "^":
			v.bzs[i].p.r = (v.bzs[i].p.r + 1) % v.h
		case "v":
			v.bzs[i].p.r = (v.bzs[i].p.r - 1 + v.h) % v.h
		}
	}
	v.sort()
}

func (v *valley) sort() {
	sort.Slice(v.bzs, func(i, j int) bool {
		if v.bzs[i].p.r < v.bzs[j].p.r {
			return true
		}
		if v.bzs[i].p.r == v.bzs[j].p.r && v.bzs[i].p.c < v.bzs[j].p.c {
			return true
		}
		return false
	})
}

type pos struct{ r, c int }

type dir struct{ r, c int }

func (p pos) move(d dir) pos {
	return pos{p.r + d.r, p.c + d.c}
}

type blizzard struct {
	p   pos
	dir string
}

func parse(lines []string) *valley {
	bzs := make([]blizzard, 0, 100)
	for r := 1; r < len(lines)-1; r++ {
		chars := u.SplitToChars(lines[r][1 : len(lines[r])-1])
		for c, ch := range chars {
			if ch != "." {
				bzs = append(bzs, blizzard{pos{r - 1, c}, ch})
			}
		}
	}
	v := valley{
		w:   len(lines[0]) - 2,
		h:   len(lines) - 2,
		bzs: bzs,
	}
	return &v
}
