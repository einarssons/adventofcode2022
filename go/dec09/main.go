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

type pair struct {
	x, y int
}

func task1(lines []string) int {
	var head, tail pair
	visited := set.FromSlice([]pair{tail})
	for _, line := range lines {
		var step pair
		dir, b, _ := strings.Cut(line, " ")
		switch dir {
		case "R":
			step = pair{1, 0}
		case "L":
			step = pair{-1, 0}
		case "U":
			step = pair{0, 1}
		case "D":
			step = pair{0, -1}
		}
		nrSteps := u.Atoi(b)
		for i := 0; i < nrSteps; i++ {
			head.x += step.x
			head.y += step.y
			tail = move(head, tail)
			visited.Add(tail)
		}
	}
	return visited.Len()
}

func task2(lines []string) int {
	knots := make([]pair, 10)
	visited := set.FromSlice([]pair{knots[9]})
	for _, line := range lines {
		var step pair
		dir, b, _ := strings.Cut(line, " ")
		switch dir {
		case "R":
			step = pair{1, 0}
		case "L":
			step = pair{-1, 0}
		case "U":
			step = pair{0, 1}
		case "D":
			step = pair{0, -1}
		}
		nrSteps := u.Atoi(b)
		for i := 0; i < nrSteps; i++ {
			knots[0].x += step.x
			knots[0].y += step.y
			for i := 1; i <= 9; i++ {
				knots[i] = move(knots[i-1], knots[i])
			}
			visited.Add(knots[9])
		}
	}
	return visited.Len()
}

func dist(a, b pair) int {
	xd := u.Abs(a.x - b.x)
	yd := u.Abs(a.y - b.y)
	return u.Max([]int{xd, yd})
}

func move(headPos, tailPos pair) pair {
	d := dist(headPos, tailPos)
	if d <= 1 {
		return tailPos
	}
	distX := headPos.x - tailPos.x
	distY := headPos.y - tailPos.y
	if u.Abs(distX) >= 1 {
		tailPos.x += u.Sign(distX)
	}
	if u.Abs(distY) >= 1 {
		tailPos.y += u.Sign(distY)
	}
	return tailPos
}
