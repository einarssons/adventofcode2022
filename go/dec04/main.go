package main

import (
	"fmt"
	"log"
	"strings"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	count := 0
	for _, line := range lines {
		p1, p2 := parseLine(line)
		if contains(p1, p2) {
			count++
		}
	}
	return count
}

func task2(lines []string) int {
	count := 0
	for _, line := range lines {
		p1, p2 := parseLine(line)
		if overlaps(p1, p2) {
			count++
		}
	}
	return count
}

type pair struct {
	start int
	end   int
}

func parseLine(line string) (p1, p2 pair) {
	l, r, ok := strings.Cut(line, ",")
	if !ok {
		log.Fatal("bad line")
	}
	p1 = parseRange(l)
	p2 = parseRange(r)
	return p1, p2
}

func parseRange(str string) pair {
	start, end := u.Cut(str, "-")
	return pair{u.Atoi(start), u.Atoi(end)}
}

func contains(p1, p2 pair) bool {
	if p1.start >= p2.start && p1.end <= p2.end {
		return true
	}
	if p2.start >= p1.start && p2.end <= p1.end {
		return true
	}
	return false
}

func overlaps(p1, p2 pair) bool {
	if p1.start >= p2.start && p1.end <= p2.end {
		return true
	}
	if p2.start >= p1.start && p2.end <= p1.end {
		return true
	}
	if p1.start >= p2.start && p1.start <= p2.end { // p1 starts in p2
		return true
	}
	if p1.end >= p2.start && p1.end <= p2.end { // p1 ends in p2
		return true
	}
	if p2.start >= p1.start && p2.start <= p1.end { // p2 starts in p1
		return true
	}
	if p2.end >= p1.start && p2.end <= p1.end { // p2 ends in p1
		return true
	}
	return false
}
