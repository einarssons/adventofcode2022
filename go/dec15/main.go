package main

import (
	"fmt"
	"log"

	"github.com/chrispappas/golang-generics-set/set"
	u "github.com/einarssons/adventofcode2022/go/utils"
	"github.com/oriser/regroup"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines, 2000000))
	fmt.Println("task2: ", task2(lines, 4000000))
}

func task1(lines []string, y int) int {
	a := newArea()
	c := newCovered(y)
	for _, line := range lines {
		sensor, beacon := ParseLine(line)
		a.addPair(sensor, beacon)
		c.addPair(sensor, beacon)
	}
	return findCovered(a, c)
}

type pair struct {
	s pos
	b pos
}

func task2(lines []string, size int) int {
	a := newArea()
	pairs := make([]pair, 0, len(lines))
	for _, line := range lines {
		sensor, beacon := ParseLine(line)
		pairs = append(pairs, pair{sensor, beacon})
		a.addPair(sensor, beacon)
	}
	for y := 0; y <= size; y++ {
		c := newCovered(y)
		for _, p := range pairs {
			c.addPair(p.s, p.b)
		}
		x, ok := findHole(a, c, size)
		if ok {
			return x*4000000 + y
		}
	}
	return -1
}

type pos struct {
	x, y int
}

func dist(a, b pos) int {
	return u.Abs(a.x-b.x) + u.Abs(a.y-b.y)
}

type area struct {
	beacons set.Set[pos]
	sensors []pos
}

func newArea() *area {
	a := area{}
	a.beacons = set.FromSlice([]pos{})
	return &a
}

func (a *area) addPair(s, b pos) {
	a.sensors = append(a.sensors, s)
	a.beacons.Add(b)
}

// Covered

type itvl struct {
	minX, maxX int
}

type covered struct {
	y     int
	itvls []itvl
}

func newCovered(y int) *covered {
	return &covered{y: y}
}

func (c *covered) addPair(s, b pos) {
	d := dist(s, b)
	yDist := u.Abs(s.y - c.y)
	if yDist <= d {
		w := d - yDist
		c.itvls = append(c.itvls, itvl{s.x - w, s.x + w})
	}
}

func (c *covered) nextCovered(x int) (int, bool) {
	m := u.MaxInt
	for _, itvl := range c.itvls {
		if itvl.minX > x && itvl.minX < m {
			m = itvl.minX
		}
	}
	if m < u.MaxInt {
		return m, true
	}
	return 0, false
}

// start in covered, find next uncovered.
func (c *covered) nextUncovered(x int) int {
	if !c.isCovered(x) {
		panic("not covered")
	}
	for {
		for _, itvl := range c.itvls {
			if itvl.minX <= x && itvl.maxX >= x {
				x = itvl.maxX + 1
			}
			ok := c.isCovered(x)
			if !ok {
				return x
			}
		}
	}
}

func (c *covered) isCovered(x int) bool {
	for _, itvl := range c.itvls {
		if itvl.minX <= x && x <= itvl.maxX {
			return true
		}
	}
	return false
}

func (c *covered) nrBeaconsInInterval(a *area, start, end int) int {
	nrBeacons := 0
	for x := start; x < end; x++ {
		for _, b := range a.beacons.Values() {
			if b.x == x && b.y == c.y {
				nrBeacons++
			}
		}
	}
	return nrBeacons
}

func findCovered(a *area, c *covered) int {
	end := -u.MaxInt
	nrCovered := 0
	for {
		start, ok := c.nextCovered(end)
		if !ok {
			break
		}
		end = c.nextUncovered(start)
		nrB := c.nrBeaconsInInterval(a, start, end)
		fmt.Println(start, end, nrB)
		nrCovered += end - start - nrB
	}
	return nrCovered
}

func findHole(a *area, c *covered, size int) (int, bool) {
	end := -u.MaxInt
	start, ok := c.nextCovered(end)
	if !ok {
		panic("bad start")
	}
	if start > 0 {
		panic("bad start 2")
	}
	end = c.nextUncovered(start)
	if end <= size {
		return end, true
	}
	return -1, false
}

// Parse data
var rex = regroup.MustCompile(`Sensor at x=(?P<sx>[\-\d]+), y=(?P<sy>[-\d]+): closest beacon is at x=(?P<bx>[-\d]+), y=(?P<by>[-\d]+)`)

type Line struct {
	Sx int `regroup:"sx"`
	Sy int `regroup:"sy"`
	Bx int `regroup:"bx"`
	By int `regroup:"by"`
}

// ParseCommand parses a "verb value" from a line.
func ParseLine(line string) (sensor, beacon pos) {
	l := Line{}
	if err := rex.MatchToTarget(line, &l); err != nil {
		log.Fatal(err)
	}
	return pos{l.Sx, l.Sy}, pos{l.Bx, l.By}
}
