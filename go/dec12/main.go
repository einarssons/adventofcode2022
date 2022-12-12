package main

import (
	"fmt"
	"sort"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

type vector struct {
	r int
	c int
}

const BIG = 1000000000

func task1(lines []string) int {
	levels := u.CreateCharGridFromLines(lines)
	pathCosts := u.CreateZeroDigitGrid(levels.Width, levels.Height)
	pathCosts.SetValue(BIG)
	origins := u.CreateEmptyCharGrid(levels.Width, levels.Height)
	origins.SetValue(".")
	startRow, startCol := levels.Find("S")
	levels.Grid[startRow][startCol] = "a"
	pathCosts.Grid[startRow][startCol] = 0
	start := vector{startRow, startCol}
	endRow, endCol := levels.Find("E")
	fmt.Printf("end: %d %d\n", endRow, endCol)
	end := vector{endRow, endCol}
	levels.Grid[end.r][end.c] = "z"
	for {
		if pathCosts.Grid[end.r][end.c] < BIG {
			break
		}
		lowestKnown := findLowestVisited(pathCosts, levels)
		for _, v := range lowestKnown {
			pos := vector{v.r, v.c}
			neighbor, ok := findLowestUnvisited(pos, pathCosts, levels)
			if !ok {
				continue
			}
			pathCosts.Grid[neighbor.r][neighbor.c] = pathCosts.Grid[v.r][v.c] + 1
			var from string
			switch {
			case neighbor.r-v.r == 1:
				from = "^"
			case neighbor.r-v.r == -1:
				from = "v"
			case neighbor.c-v.c == 1:
				from = "<"
			case neighbor.c-v.c == -1:
				from = ">"
			default:
				panic("bad")
			}
			origins.Grid[neighbor.r][neighbor.c] = from
		}
	}
	fmt.Printf("end path cost: %d\n", pathCosts.Grid[end.r][end.c])
	fmt.Println(origins.String())
	pathGrid := makePathGrid(end, start, origins)
	fmt.Println(pathGrid.String())
	l := pathLength(end, start, origins)
	return l
}

// 420 is too high
// Go backwards insted until level a has been found.
// This way we'll only find connectedpaths that goes to end
func task2(lines []string) int {
	levels := u.CreateCharGridFromLines(lines)
	endRow, endCol := levels.Find("E")
	fmt.Printf("end: %d %d\n", endRow, endCol)
	end := vector{endRow, endCol}
	levels.Grid[end.r][end.c] = "z"
	r, c := levels.Find("S")
	levels.Grid[r][c] = "a"
	pathCosts := u.CreateZeroDigitGrid(levels.Width, levels.Height)
	pathCosts.SetValue(BIG)
	pathCosts.Grid[end.r][end.c] = 0
	shortestPath := BIG
	var shortestStart vector
	for {
		cv, ok := cheapestVisitedUnfinished(pathCosts, levels)
		if !ok {
			break
		}
		neighbor, ok := findCheapestNeighbor(cv, pathCosts, levels)
		if !ok {
			continue
		}
		neighborPathCost := pathCosts.Grid[cv.r][cv.c] + 1
		pathCosts.Grid[neighbor.r][neighbor.c] = neighborPathCost
		if neighborPathCost < shortestPath {
			if levels.Grid[neighbor.r][neighbor.c] == "a" {
				shortestPath = neighborPathCost
				shortestStart = neighbor
			}
		}
	}
	fmt.Printf("start point: (%d, %d)\n", shortestStart.r, shortestStart.c)
	return shortestPath
}

// findCheapestVisited returns the node with the lowest path, that has
// that does not have minimal value, but accessible unexplored neighbors.
// They should also have a neighbor that has not been visited.
func findCheapestVisited(pathCosts u.DigitGrid, levels u.CharGrid) []pathCost {
	positions := make([]pathCost, 0, pathCosts.Width*pathCosts.Height)
	for r := 0; r < pathCosts.Height; r++ {
	columtLoop:
		for c := 0; c < pathCosts.Width; c++ {
			cost := pathCosts.Grid[r][c]
			if cost == BIG {
				continue // This node is not visisted
			}
			for i := 0; i < 4; i++ {
				nr, nc := r, c
				switch i {
				case 0:
					nr++
				case 1:
					nr--
				case 2:
					nc++
				case 3:
					nc--
				}
				if !pathCosts.InBounds(nr, nc) {
					continue // Neighbor outside grid
				}
				neighborLevel := u.FirstAsciiNr(levels.Grid[nr][nc])
				if neighborLevel <= u.FirstAsciiNr(levels.Grid[r][c])+1 {
					positions = append(positions, pathCost{r, c, cost})
					continue columtLoop
				}
			}
		}
	}
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].cost < positions[j].cost
	})
	return positions
}

// findCheapestNeighbor finds a neighbor which is not yet visited and can be visited.
func findCheapestNeighbor(pos vector, pathCosts u.DigitGrid, levels u.CharGrid) (vector, bool) {
	level := u.FirstAsciiNr(levels.Grid[pos.r][pos.c])
	for i := 0; i < 4; i++ {
		nr, nc := pos.r, pos.c
		switch i {
		case 0:
			nr++
		case 1:
			nr--
		case 2:
			nc++
		case 3:
			nc--
		}
		if !pathCosts.InBounds(nr, nc) {
			continue // Neighbor outside grid
		}
		if pathCosts.Grid[nr][nc] != BIG {
			continue // Already visited
		}
		neighborLevel := u.FirstAsciiNr(levels.Grid[nr][nc])
		if neighborLevel >= level-1 {
			return vector{nr, nc}, true
		}
	}
	return pos, false
}

// cheapestVisitedUnfinished should have a neighbor from which we could have come.
// It should also not alraedy have level "a"
func cheapestVisitedUnfinished(pathCosts u.DigitGrid, levels u.CharGrid) (vector, bool) {
	minCost := BIG
	var minPos vector
	for r := 0; r < pathCosts.Height; r++ {
		for c := 0; c < pathCosts.Width; c++ {
			thisPathCost := pathCosts.Grid[r][c]
			if thisPathCost == BIG {
				continue // not yet visited
			}
			if levels.Grid[r][c] == "a" {
				continue // Already visited and lowest level
			}
			thisLevel := u.FirstAsciiNr(levels.Grid[r][c])
			for i := 0; i < 4; i++ {
				nr, nc := r, c
				switch i {
				case 0:
					nr++
				case 1:
					nr--
				case 2:
					nc++
				case 3:
					nc--
				}
				if !pathCosts.InBounds(nr, nc) {
					continue // Neighbor outside grid
				}
				if pathCosts.Grid[nr][nc] != BIG {
					continue // Already visited
				}
				neighborLevel := u.FirstAsciiNr(levels.Grid[nr][nc])
				if neighborLevel >= thisLevel-1 {
					if thisPathCost < minCost {
						minCost = thisPathCost
						minPos = vector{r, c}
					}
				}
			}
		}
	}
	if minCost == BIG {
		return vector{0, 0}, false
	}
	return minPos, true
}

type pathCost struct {
	r, c, cost int
}

func findLowestPositions(levels u.CharGrid) []vector {
	var p []vector
	for r := 0; r < levels.Height; r++ {
		for c := 0; c < levels.Width; c++ {
			if levels.Grid[r][c] == "a" {
				p = append(p, vector{r, c})
			}
		}
	}
	return p
}

// findLowestVisited returns slice of coordinates for lowest visited places.
// They should also have a neighbor that has not been visited.
func findLowestVisited(pathCosts u.DigitGrid, levels u.CharGrid) []pathCost {
	positions := make([]pathCost, 0, pathCosts.Width*pathCosts.Height)
	for r := 0; r < pathCosts.Height; r++ {
	columtLoop:
		for c := 0; c < pathCosts.Width; c++ {
			cost := pathCosts.Grid[r][c]
			if cost == BIG {
				continue // This node is not visisted
			}
			for i := 0; i < 4; i++ {
				nr, nc := r, c
				switch i {
				case 0:
					nr++
				case 1:
					nr--
				case 2:
					nc++
				case 3:
					nc--
				}
				if !pathCosts.InBounds(nr, nc) {
					continue // Neighbor outside grid
				}
				neighborLevel := u.FirstAsciiNr(levels.Grid[nr][nc])
				if neighborLevel <= u.FirstAsciiNr(levels.Grid[r][c])+1 {
					positions = append(positions, pathCost{r, c, cost})
					continue columtLoop
				}
			}

		}
	}
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].cost < positions[j].cost
	})
	return positions
}

// findLowestUnvisited finds a neighbor to the lowest visited if possible
func findLowestUnvisited(lowestVisited vector, pathCosts u.DigitGrid, levels u.CharGrid) (vector, bool) {
	level := u.FirstAsciiNr(levels.Grid[lowestVisited.r][lowestVisited.c])
	for i := 0; i < 4; i++ {
		neighbor := lowestVisited
		switch i {
		case 0:
			neighbor.r--
		case 1:
			neighbor.r++
		case 2:
			neighbor.c--
		case 3:
			neighbor.c++
		}
		if !pathCosts.InBounds(neighbor.r, neighbor.c) {
			continue
		}
		if pathCosts.Grid[neighbor.r][neighbor.c] != BIG {
			continue // already visited
		}
		if u.FirstAsciiNr(levels.Grid[neighbor.r][neighbor.c]) <= level+1 {
			return neighbor, true
		}
	}
	return vector{0, 0}, false
}

func lowest(pos vector, pathCosts u.DigitGrid, levels, origins u.CharGrid) (from vector, ok bool, cost int) {
	level := levels.Grid[pos.r][pos.c]
	bestNeighbor := vector{0, 0}
	bestNeighborCost := BIG
	for r := pos.r - 1; r <= pos.r+1; r++ {
		for c := pos.c - 1; c <= pos.c+1; c++ {
			if r < 0 || c < 0 || r >= pathCosts.Height || c >= pathCosts.Width {
				continue
			}
			neighborPathCost := pathCosts.Grid[r][c]
			if neighborPathCost == BIG {
				continue
			}
			if (r-pos.r)*(c-pos.c) != 0 { // Must be same row or col
				continue
			}
			neighborLevel := levels.Grid[r][c]
			levelDiff := u.FirstAsciiNr(level) - u.FirstAsciiNr(neighborLevel)
			if levelDiff <= 1 {
				if neighborPathCost < bestNeighborCost {
					bestNeighbor = vector{r, c}
					bestNeighborCost = neighborPathCost
				}
			}
		}
	}
	if bestNeighborCost == BIG {
		return vector{0, 0}, false, BIG
	}
	return bestNeighbor, true, bestNeighborCost + 1
}

func pathLength(end, start vector, origins u.CharGrid) int {
	pos := end
	length := 0
	for {
		switch origins.Grid[pos.r][pos.c] {
		case "^":
			pos.r--
		case "v":
			pos.r++
		case ">":
			pos.c++
		case "<":
			pos.c--
		}
		length++
		if pos == start {
			return length
		}
	}
}

// Traverse origins backwards to find path from start to end.
func makePathGrid(end, start vector, origins u.CharGrid) u.CharGrid {
	g := u.CreateEmptyCharGrid(origins.Width, origins.Height)
	g.SetValue(".")
	g.Grid[end.r][end.c] = "E"
	pos := end
	for {
		var reverse string
		switch origins.Grid[pos.r][pos.c] {
		case "<":
			pos.c--
			reverse = ">"
		case ">":
			pos.c++
			reverse = "<"
		case "^":
			pos.r--
			reverse = "v"
		case "v":
			pos.r++
			reverse = "^"
		default:
			panic("bad")
		}
		g.Grid[pos.r][pos.c] = reverse
		if pos == start {
			return g
		}
	}
}
