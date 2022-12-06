package main

import (
	"fmt"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", findPos(lines[0], 4))
	fmt.Println("task2: ", findPos(lines[0], 14))
}

func findPos(line string, nrDiffChars int) int {
	for i := nrDiffChars; i < len(line); i++ {
		s := u.CreateSet[string]()
		for j := i - (nrDiffChars - 1); j <= i; j++ {
			s.Add(string(line[j]))
		}
		if s.Size() == nrDiffChars {
			fmt.Println(s.Values())
			return i + 1
		}
	}
	return -1
}
