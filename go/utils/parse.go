package utils

import (
	"log"
	"regexp"
	"strconv"

	"github.com/oriser/regroup"
)

// SplitToInts finds all ints in line (including sign).
func SplitToInts(line string) []int {
	re := regexp.MustCompile("-?[0-9]+")
	parts := re.FindAllString(line, -1)
	var numbers []int
	for _, p := range parts {
		number, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

// SplitToChars splits a line into chars.
func SplitToChars(line string) []string {
	chars := make([]string, len(line))
	for i := 0; i < len(line); i++ {
		chars[i] = line[i : i+1]
	}
	return chars
}

func ContainsInt(x int, entries []int) bool {
	for _, n := range entries {
		if x == n {
			return true
		}
	}
	return false
}

func CountInts(entries []int) map[int]int {
	m := make(map[int]int)
	for _, n := range entries {
		m[n]++
	}
	return m
}

func CountStrings(entries []string) map[string]int {
	m := make(map[string]int)
	for _, n := range entries {
		m[n]++
	}
	return m
}

func ContainsString(x string, entries []string) bool {
	for _, n := range entries {
		if x == n {
			return true
		}
	}
	return false
}

var rex = regroup.MustCompile(`(?P<verb>[a-zA-Z]+)\s+(?P<value>\d+)`)

type Command struct {
	Verb  string `regroup:"verb"`
	Value int    `regroup:"value"`
}

// ParseCommand parses a "verb value" from a line.
func ParseCommand(line string) Command {
	c := Command{}
	if err := rex.MatchToTarget(line, &c); err != nil {
		log.Fatal(err)
	}
	return c
}
