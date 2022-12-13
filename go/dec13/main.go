package main

import (
	"fmt"
	"sort"
	"strconv"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	var idx, idxSum int
	var lStr, rStr string
	for i, line := range lines {
		switch i % 3 {
		case 0:
			lStr = line
		case 1:
			rStr = line
			idx = i/3 + 1
			order := compare(lStr, rStr)
			if order == -1 {
				idxSum += idx
			}
		case 2:
			//
		}
	}
	return idxSum
}

// 1-based index
func findIndex(line string, lines []string) int {
	for i, l := range lines {
		if line == l {
			return i + 1
		}
	}
	return 0
}

func task2(allLines []string) int {
	var lines []string
	for _, l := range allLines {
		if l != "" {
			lines = append(lines, l)
		}
	}
	divider1 := "[[2]]"
	divider2 := "[[6]]"
	lines = append(lines, divider1)
	lines = append(lines, divider2)

	sort.Slice(lines, func(i, j int) bool {
		return compare(lines[i], lines[j]) < 0
	})
	/*
		for _, l := range lines {
			fmt.Println(l)
		}
	*/
	index1 := findIndex(divider1, lines)
	index2 := findIndex(divider2, lines)
	return index1 * index2

}

// elem is either a nr >= 0, or a string to be parsed.
type elem struct {
	nr      int
	listStr string
}

// toNumber converts e to a number if possible.
func (e elem) toNumber() elem {
	if e.nr < 1 {
		i, err := strconv.Atoi(e.listStr)
		if err == nil {
			return elem{i, ""}
		}
	}
	return e
}

func parseStr(str string) []elem {
	var els []elem
	nr, err := strconv.Atoi(str)
	if err == nil {
		return []elem{{nr, ""}}
	}
	level := 0
	var lev1Start int
	var parts []string
	for i := 0; i < len(str); i++ {
		switch c := string(str[i]); c {
		case "[":
			level++
			if level == 1 {
				lev1Start = i + 1
			}
		case "]":
			level--
			if level == 0 {
				parts = append(parts, str[lev1Start:i])
			}
		case ",":
			if level == 1 {
				parts = append(parts, str[lev1Start:i])
				lev1Start = i + 1
			}
		}
	}
	for _, p := range parts {
		els = append(els, elem{-1, p})
	}
	return els
}

func compare(left, right string) int {
	lEls := parseStr(left)
	rEls := parseStr(right)
	minElems := u.Min([]int{len(lEls), len(rEls)})
	for i := 0; i < minElems; i++ {
		lEl := lEls[i].toNumber()
		rEl := rEls[i].toNumber()
		lStr := lEl.listStr
		rStr := rEl.listStr
		switch {
		case lEl.nr >= 0 && rEl.nr >= 0:
			r := u.Cmp(lEl.nr, rEl.nr)
			if r != 0 {
				if r < 0 {
					return -1
				}
				return 1
			}
			continue
		case lEl.nr >= 0:
			lStr = fmt.Sprintf("[%d]", lEl.nr)
		case rEl.nr >= 0:
			rStr = fmt.Sprintf("[%d]", rEl.nr)
		default:
			//
		}
		r := compare(lStr, rStr)
		if r != 0 {
			return r
		}
	}
	return u.Cmp(len(lEls), len(rEls))
}
