package main

import (
	"fmt"
	"strings"

	"github.com/chrispappas/golang-generics-set/set"
	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	cave := newCave()
	parseLines(lines, cave)
	nrSand := 0
	for {
		full := cave.dropSand()
		if full {
			break
		}
		nrSand++

	}
	return nrSand
}

func task2(lines []string) int {
	return 0
}

type cave struct {
	rocks  set.Set[string]
	sand   set.Set[string]
	bottom int // lowest point
}

type pos struct {
	x, y int
}

func (p pos) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func toPos(s string) pos {
	s = strings.TrimSpace(s)
	a, b := u.Cut(s, ",")
	return pos{u.Atoi(a), u.Atoi(b)}
}

func newCave() *cave {
	c := cave{}
	c.rocks = set.FromSlice([]string{})
	c.sand = set.FromSlice([]string{})
	c.bottom = 0
	return &c
}

func (c *cave) filled(p pos) bool {
	str := p.String()
	return c.sand.Has(str) || c.rocks.Has(str)
}

func (c *cave) addRock(p pos) {
	c.rocks.Add(p.String())
	if p.y > c.bottom {
		c.bottom = p.y
	}
}

func (c *cave) addSand(p pos) {
	c.rocks.Add(p.String())
}

// dropSand returns true if beyond edge
func (c *cave) dropSand() bool {
	p := pos{500, 0}
	for {
		tryP := p
		if p.y > c.bottom {
			return true
		}
		p.y++ // down
		if !c.filled(p) {
			continue
		}
		p.x-- // left
		if !c.filled(p) {
			continue
		}
		p.x += 2 // right
		if !c.filled(p) {
			continue
		}
		// cannot move
		c.addSand(tryP)
		return false
	}
}

func (c *cave) addWall(p1, p2 pos) {
	if p1 == p2 {
		c.addRock(p1)
		return
	}
	if p1.x == p2.x {
		if p1.y < p2.y {
			for y := p1.y; y <= p2.y; y++ {
				c.addRock(pos{p1.x, y})
			}
		} else {
			for y := p2.y; y <= p1.y; y++ {
				c.addRock(pos{p1.x, y})
			}
		}
		return
	}
	if p1.x < p2.x {
		for x := p1.x; x <= p2.x; x++ {
			c.addRock(pos{x, p1.y})
		}
	} else {
		for x := p2.x; x <= p1.x; x++ {
			c.addRock(pos{x, p1.y})
		}
	}
}

func parseLines(lines []string, c *cave) {
	for _, l := range lines {
		parts := strings.Split(l, " -> ")
		for i := 0; i < len(parts)-1; i++ {
			pos1 := toPos(parts[i])
			pos2 := toPos(parts[i+1])
			c.addWall(pos1, pos2)
		}
	}
}
