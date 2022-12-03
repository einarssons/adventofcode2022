package main

import (
	"fmt"

	"github.com/chrispappas/golang-generics-set/set"
	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	totPoints := 0
	for _, line := range lines {
		parts := u.SplitToChars(line)
		left := parts[:len(parts)/2]
		right := parts[len(parts)/2:]
		leftSet := set.FromSlice(left)
		rightSet := set.FromSlice(right)
		common := leftSet.Intersection(rightSet).Values()
		totPoints += summedValues(common)
	}

	return totPoints
}

func summedValues(s []string) int {
	sum := 0
	for _, c := range s {
		sum += value(c)
	}
	return sum
}

func value(v string) int {
	nr := u.FirstAsciiNr(v)
	aNr := u.FirstAsciiNr("a")
	ANr := u.FirstAsciiNr("A")
	switch {
	case nr > aNr:
		return nr - aNr + 1
	default:
		return nr - ANr + 27
	}
}

func task2(lines []string) int {
	totPoints := 0
	var sets [3]set.Set[string]
	for i, line := range lines {
		imod := i % 3
		parts := u.SplitToChars(line)
		s := set.FromSlice(parts)
		sets[imod] = s
		if i%3 == 2 {
			common := sets[0].Intersection(sets[1]).Intersection(sets[2]).Values()
			totPoints += summedValues(common)
		}
	}
	return totPoints
}
