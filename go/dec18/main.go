package main

import (
	"fmt"

	"github.com/chrispappas/golang-generics-set/set"
	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	var cubes []cube
	for _, line := range lines {
		ints := u.SplitToInts(line)
		cubes = append(cubes, cube{ints[0], ints[1], ints[2]})
	}
	totalArea := 0
	for _, c := range cubes {
		totalArea += 6 - coveredSides(c, cubes)
	}
	return totalArea
}

func task2(lines []string) int {
	var cubes []cube
	for _, line := range lines {
		ints := u.SplitToInts(line)
		cubes = append(cubes, cube{ints[0], ints[1], ints[2]})
	}
	min := dist{u.MaxInt, u.MaxInt, u.MaxInt}
	max := dist{-u.MaxInt, -u.MaxInt, -u.MaxInt}

	for _, c := range cubes {
		for i := 0; i < 3; i++ {
			if c[i] < min[i] {
				min[i] = c[i]
			}
			if c[i] > max[i] {
				max[i] = c[i]
			}
		}
	}
	tuples := make([]tuple, 0, len(cubes))
	for _, c := range cubes {
		tuples = append(tuples, tuple{c[0], c[1], c[2]})
	}
	lava := set.FromSlice(tuples)
	dt := newDroplet(min, max)

	dt.fillWater(lava)
	surface := dt.surface()
	return surface
}

type droplet struct {
	water [][][]bool
	min   dist
	max   dist
}

func newDroplet(min, max dist) *droplet {
	w := make([][][]bool, 0, max[0]-min[0]+1)
	for x := min[0]; x <= max[0]; x++ {
		yw := make([][]bool, 0, max[1]-min[1]+1)
		for y := min[1]; y <= max[1]; y++ {
			yw = append(yw, make([]bool, max[2]-min[2]+1))
		}
		w = append(w, yw)
	}
	dt := droplet{
		water: w,
		min:   min,
		max:   max,
	}
	return &dt
}

func (dt *droplet) setW(x, y, z int) {
	dt.water[x-dt.min[0]][y-dt.min[1]][z-dt.min[2]] = true
}

func (dt *droplet) isW(x, y, z int) bool {
	return dt.water[x-dt.min[0]][y-dt.min[1]][z-dt.min[2]]
}

func (dt *droplet) nrWater() int {
	nr := 0
	for x := dt.min[0]; x <= dt.max[0]; x++ {
		for y := dt.min[1]; y <= dt.max[1]; y++ {
			for z := dt.min[2]; z <= dt.max[2]; z++ {
				if dt.isW(x, y, z) {
					nr++
				}
			}
		}
	}
	return nr
}

func (dt *droplet) isEdge(x, y, z int) bool {
	low := x == dt.min[0] || y == dt.min[1] || z == dt.min[2]
	high := x == dt.max[0] || y == dt.max[1] || z == dt.max[2]
	return low || high
}

func (dt *droplet) isOut(x, y, z int) bool {
	if x < dt.min[0] || y < dt.min[1] || z < dt.min[2] {
		return true
	}
	if x > dt.max[0] || y > dt.max[1] || z > dt.max[2] {
		return true
	}
	return false
}

func (dt *droplet) volume() int {
	f := 1
	for i := 0; i < 3; i++ {
		f *= dt.max[i] - dt.min[i] + 1
	}
	return f
}

func (dt *droplet) fillWater(lava set.Set[tuple]) {
	for x := dt.min[0]; x <= dt.max[0]; x++ {
		for y := dt.min[1]; y <= dt.max[1]; y++ {
			for z := dt.min[2]; z <= dt.max[2]; z++ {
				t := tuple{x, y, z}
				if !lava.Has(t) {
					if dt.isEdge(x, y, z) {
						dt.setW(x, y, z)
					}
				}
			}
		}
	}
	nrWater := dt.nrWater()
	fmt.Printf("initial water: %d (%d)\n", nrWater, dt.volume())
	//dt.print(true)
	//dt.print(false)
	for {
		nrNewWater := dt.fillMore(lava)
		if nrNewWater == 0 {
			break
		}
		//dt.print(false)
		fmt.Printf("new water: %d\n", nrNewWater)
	}
	nrWater = dt.nrWater()
	fmt.Printf("final water: %d (%d)\n", nrWater, dt.volume())
}

var neighbors = [6][3]int{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}}

func (dt *droplet) fillMore(lava set.Set[tuple]) int {
	nrNewWater := 0
	for x := dt.min[0] + 1; x <= dt.max[0]-1; x++ {
		for y := dt.min[1] + 1; y <= dt.max[1]-1; y++ {
			for z := dt.min[2] + 1; z <= dt.max[2]-1; z++ {
				for _, n := range neighbors {
					if !dt.isW(x, y, z) && dt.isW(x+n[0], y+n[1], z+n[2]) {
						if !lava.Has(tuple{x, y, z}) {
							dt.setW(x, y, z)
							nrNewWater++
						}
					}
				}
			}
		}
	}
	return nrNewWater
}

func (dt *droplet) surface() int {
	surface := 0
	for x := dt.min[0]; x <= dt.max[0]; x++ {
		for y := dt.min[1]; y <= dt.max[1]; y++ {
			for z := dt.min[2]; z <= dt.max[2]; z++ {
				for _, n := range neighbors {
					if !dt.isW(x, y, z) && (dt.isOut(x+n[0], y+n[1], z+n[2]) || dt.isW(x+n[0], y+n[1], z+n[2])) {
						surface += 1
					}
				}
			}
		}
	}
	return surface
}

func (dt *droplet) print(water bool) {
	if water {
		fmt.Println("\nwater:")
	} else {
		fmt.Println("\nNot water:")
	}
	count := 0
	for x := dt.min[0]; x <= dt.max[0]; x++ {
		for y := dt.min[1]; y <= dt.max[1]; y++ {
			for z := dt.min[2]; z <= dt.max[2]; z++ {
				if dt.isW(x, y, z) == water {
					fmt.Printf("(%d, %d, %d)\n", x, y, z)
					count++
				}
			}
		}
	}
	fmt.Printf("nr: %d\n", count)
	fmt.Println()
}

type cube [3]int
type dist [3]int

func absDist(a, b cube) int {
	return u.Abs(a[0]-b[0]) + u.Abs(a[1]-b[1]) + u.Abs(a[2]-b[2])
}

func distance(a, b cube) dist {
	var d dist
	for i := 0; i < 3; i++ {
		d[i] = a[i] - b[i]
	}
	return d
}

type tuple struct {
	x, y, z int
}

func tupleFromDist(d dist) tuple {
	return tuple{d[0], d[1], d[2]}
}

func coveredSides(c cube, cs []cube) int {
	sides := set.FromSlice([]tuple{})
	for _, b := range cs {
		if absDist(c, b) == 1 {
			sides.Add(tupleFromDist(distance(c, b)))
		}
	}
	return sides.Len()
}
