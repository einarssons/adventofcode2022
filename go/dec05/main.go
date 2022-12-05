package main

import (
	"fmt"
	"log"
	"strings"

	u "github.com/einarssons/adventofcode2022/go/utils"
	"github.com/oriser/regroup"
)

func main() {
	lines := u.ReadRawLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

var rex = regroup.MustCompile(`move (?P<nr>\d+) from (?P<from>\d+) to (?P<to>\d+)`)

type Move struct {
	Nr   int `regroup:"nr"`
	From int `regroup:"from"`
	To   int `regroup:"to"`
}

// ParseMove parses a move command from a line.
func ParseMove(line string) Move {
	m := Move{}
	if err := rex.MatchToTarget(line, &m); err != nil {
		log.Fatal(err)
	}
	return m
}

func task1(lines []string) string {
	l := len(lines[0]) + 1
	nrStacks := l / 4
	stacks := make([]u.StackStrings, nrStacks)
	state := "cargo"
	for _, line := range lines {
		switch state {
		case "cargo":
			found := parseStacks(line, stacks)
			if found {
				continue
			}
			// Line of numbers. Ready with initial state.
			// Reverse stacks since we read from the top.
			for _, s := range stacks {
				s.Reverse()
			}
			state = "blank"
		case "blank":
			// Blank separator line
			state = "moves"
		case "moves":
			// Move cargo by pop from one stack and push to another.
			m := ParseMove(line)
			for i := 0; i < m.Nr; i++ {
				item, _ := stacks[m.From-1].Pop()
				stacks[m.To-1].Push(item)
			}
		}
	}
	// Find message of top stacks.
	msg := ""
	for _, s := range stacks {
		item, _ := s.Pop()
		msg += item
	}
	return msg
}

func task2(lines []string) string {
	l := len(lines[0]) + 1
	nrStacks := l / 4
	stacks := make([]u.StackStrings, nrStacks)
	state := "cargo"
	for _, line := range lines {
		switch state {
		case "cargo":
			found := parseStacks(line, stacks)
			if found {
				continue
			}
			// Line of numbers. Ready with initial state.
			// Reverse stacks since we read from the top.
			for _, s := range stacks {
				s.Reverse()
			}
			state = "blank"
		case "blank":
			// Blank separator line
			state = "moves"
		case "moves":
			// Move multiple cargos. Similuate using intermediate stack.
			m := ParseMove(line)
			localStack := u.StackStrings{}
			for i := 0; i < m.Nr; i++ {
				item, _ := stacks[m.From-1].Pop()
				localStack.Push(item)
			}
			for i := 0; i < m.Nr; i++ {
				item, _ := localStack.Pop()
				stacks[m.To-1].Push(item)
			}
		}
	}
	// Find message of top stacks.
	msg := ""
	for _, s := range stacks {
		item, _ := s.Pop()
		msg += item
	}
	return msg
}

// parseStacks splits a line into parts and look for [X] cargo.
// The cargo is added to stacks depending on index.
func parseStacks(line string, stacks []u.StackStrings) bool {
	nrStacks := len(stacks)
	for i := 0; i < nrStacks; i++ {
		end := 4 * (i + 1)
		if end > len(line) {
			end = len(line)
		}
		sChars := string(line[4*i : end])
		part := strings.Trim(sChars, " ")
		if part == "" {
			continue
		}
		if !strings.HasPrefix(sChars, "[") {
			return false
		}
		c := string(part[1])
		stacks[i].Push(c)
	}
	return true
}
