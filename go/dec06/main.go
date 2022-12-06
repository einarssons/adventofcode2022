package main

import (
	"fmt"

	"github.com/chrispappas/golang-generics-set/set"
	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", findPos(lines[0], 4))
	fmt.Println("task2: ", findPos(lines[0], 14))
}

func findPos(line string, nrDiffChars int) int {
	for i := nrDiffChars - 1; i < len(line); i++ {
		s := set.FromSlice(u.SplitToChars(line[i-nrDiffChars+1 : i+1]))
		if s.Len() == nrDiffChars {
			return i + 1
		}
	}
	return -1
}
