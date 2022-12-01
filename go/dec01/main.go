package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	task1(lines)
	task2(lines)
}

func task1(lines []string) {
	totals := make([]int, 0, 100)
	currTotal := 0
	for _, line := range lines {
		if line == "" {
			totals = append(totals, currTotal)
			currTotal = 0
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		currTotal += val
	}
	fmt.Println(u.Max(totals))
}

func task2(lines []string) {
	totals := make([]int, 0, 100)
	currTotal := 0
	for _, line := range lines {
		if line == "" {
			totals = append(totals, currTotal)
			currTotal = 0
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		currTotal += val
	}
	totals = append(totals, currTotal)
	max := 0
	sort.Ints(totals)
	l := len(totals)
	for i := l - 3; i <= l-1; i++ {
		max += totals[i]
	}
	fmt.Println(max)
}
