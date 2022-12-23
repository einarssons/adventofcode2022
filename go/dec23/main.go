package main

import (
	"fmt"
	"sort"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	elves := parseElves(lines)
	initDirs()
	for round := 1; round <= 10; round++ {
		proposedMoves := make(map[pos]int) // Count moves to find doubles
		for elfNr, e := range elves {
			m, ok := e.proposeMove(elves, elfNr)
			elves[elfNr].pMove = m
			if ok {
				proposedMoves[m] += 1
			}
		}
		for i := range elves {
			elves[i].move(proposedMoves)
		}
		updateInitDirs()
		sortElves(elves)
	}
	printElves(elves)
	ul, br := findBounds(elves)
	area := (br.r - ul.r + 1) * (br.c - ul.c + 1)
	nrEmpty := area - len(elves)
	return nrEmpty
}

func printElves(elves []elf) {
	ul, br := findBounds(elves)
	eNr := 0
	for r := ul.r; r <= br.r; r++ {
		for c := ul.c; c <= br.c; c++ {
			p := pos{r, c}
			if eNr >= len(elves) {
				fmt.Printf(".")
				continue
			}
			if elves[eNr].p == p {
				fmt.Printf("#")
				eNr++
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func sortElves(elves []elf) {
	sort.Slice(elves, func(i, j int) bool {
		if elves[i].p.r < elves[j].p.r {
			return true
		}
		if elves[i].p.r == elves[j].p.r && elves[i].p.c < elves[j].p.c {
			return true
		}
		return false
	})
}

func task2(lines []string) int {
	elves := parseElves(lines)
	initDirs()
	for round := 1; round < u.MaxInt; round++ {
		proposedMoves := make(map[pos]int) // Count moves to find doubles
		for elfNr, e := range elves {
			m, ok := e.proposeMove(elves, elfNr)
			elves[elfNr].pMove = m
			if ok {
				proposedMoves[m] += 1
			}
		}
		if len(proposedMoves) == 0 {
			printElves(elves)
			return round
		}
		for i := range elves {
			elves[i].move(proposedMoves)
		}
		updateInitDirs()
		sortElves(elves)
	}
	return -1
}

type pos struct {
	r, c int
}

func (p pos) move(d dir) pos {
	return pos{p.r + d.r, p.c + d.c}
}

type dir struct {
	r, c int
}

type elf struct {
	p     pos
	pMove pos
}

var NW = dir{-1, -1}
var N = dir{-1, 0}
var NE = dir{-1, 1}
var W = dir{0, -1}
var E = dir{0, 1}
var SW = dir{1, -1}
var S = dir{1, 0}
var SE = dir{1, 1}

var dirs [][]dir
var mDirs []dir

func initDirs() {
	dirs = make([][]dir, 4)
	dirs[0] = []dir{N, NE, NW}
	dirs[1] = []dir{S, SE, SW}
	dirs[2] = []dir{W, NW, SW}
	dirs[3] = []dir{E, NE, SE}
	mDirs = make([]dir, 4)
	mDirs[0] = N
	mDirs[1] = S
	mDirs[2] = W
	mDirs[3] = E
}

func updateInitDirs() {
	d := dirs[0]
	copy(dirs, dirs[1:])
	dirs[3] = d
	md := mDirs[0]
	copy(mDirs, mDirs[1:])
	mDirs[3] = md
}

// proposeMove proposes a move if true
func (e *elf) proposeMove(elves []elf, elfNr int) (pos, bool) {
	nbs := findNeighbors(elfNr, elves)
	if len(nbs) == 0 {
		return e.p, false
	}
	for i, ds := range dirs {
		if e.checkDirs(ds, nbs) {
			pMove := e.p.move(mDirs[i])
			return pMove, true
		}
	}
	return e.p, false
}

func (e *elf) move(proposedMoves map[pos]int) {
	if e.pMove != e.p && proposedMoves[e.pMove] == 1 {
		e.p = e.pMove
	}
}

func findNeighbors(elfNr int, elves []elf) []pos {
	nbs := make([]pos, 0, 4)
	elf := elves[elfNr]
	eRow := elf.p.r
	eCol := elf.p.c
	nrElves := len(elves)
	for nr := elfNr - 1; nr >= 0; nr-- {
		nbElf := elves[nr]
		nbr := nbElf.p.r
		nbc := nbElf.p.c
		if nbr < eRow-1 {
			break
		}
		if nbc < eCol-1 || nbc > eCol+1 {
			continue
		}
		nbs = append(nbs, nbElf.p)
	}
	for nr := elfNr + 1; nr < nrElves; nr++ {
		nbElf := elves[nr]
		nbr := nbElf.p.r
		nbc := nbElf.p.c
		if nbr > eRow+1 {
			break
		}
		if nbc < eCol-1 || nbc > eCol+1 {
			continue
		}
		nbs = append(nbs, nbElf.p)
	}
	return nbs
}

func findBounds(elves []elf) (topLeft pos, bottomRight pos) {
	topLeft = pos{u.MaxInt, u.MaxInt}
	bottomRight = pos{-u.MaxInt, -u.MaxInt}
	for _, e := range elves {
		if e.p.c < topLeft.c {
			topLeft.c = e.p.c
		}
		if e.p.r < topLeft.r {
			topLeft.r = e.p.r
		}
		if e.p.c > bottomRight.c {
			bottomRight.c = e.p.c
		}
		if e.p.r > bottomRight.r {
			bottomRight.r = e.p.r
		}
	}
	return topLeft, bottomRight
}

// checkDirs returns true if no other elf in the directions
func (e *elf) checkDirs(ds []dir, neighborPos []pos) bool {
	for _, nb := range neighborPos {
		for _, d := range ds {
			np := pos{e.p.r + d.r, e.p.c + d.c}
			if np == nb {
				return false
			}
		}
	}
	return true
}

func parseElves(lines []string) []elf {
	var elves []elf
	for row, line := range lines {
		for col, c := range u.SplitToChars(line) {
			if c == "#" {
				elves = append(elves, elf{p: pos{row, col}})
			}
		}
	}
	return elves
}
