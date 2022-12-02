package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	fmt.Printf("task1: %d\n", task1(lines))
	fmt.Printf("task2: %d\n", task2(lines))
}

func task1(lines []string) int {
	totalPoints := 0
	for _, line := range lines {
		a, b, _ := strings.Cut(line, " ")
		win := 0
		form := 0
		switch b {
		case "X": // Rock
			form = 1
			switch a {
			case "A": // Rock
				win = 3
			case "B": // Paper
				win = 0
			case "C": // Scissors
				win = 6
			}
		case "Y": // Paper
			form = 2
			switch a {
			case "A": // Rock
				win = 6
			case "B": // Paper
				win = 3
			case "C": // Scissors
				win = 0
			}
		case "Z": // Scissors
			form = 3
			switch a {
			case "A": // Rock
				win = 0
			case "B": // Paper
				win = 6
			case "C": // Scissors
				win = 3
			}
		}
		totalPoints += win + form
	}
	return totalPoints
}

func task2(lines []string) int {
	totalPoints := 0
	for _, line := range lines {
		a, b, _ := strings.Cut(line, " ")
		win := 0
		form := 0
		switch a {
		case "A": // Rock
			switch b {
			case "X": // lose
				win = 0
				form = 3 // Scissors
			case "Y": // draw
				win = 3
				form = 1 // Rock
			case "Z": // win
				win = 6
				form = 2 // Paper
			}
		case "B": // Paper
			switch b {
			case "X": // lose
				win = 0
				form = 1 // Rock
			case "Y": // draw
				win = 3
				form = 2 // Paper
			case "Z": // win
				win = 6
				form = 3 // Scissors
			}
		case "C": // Scissors
			switch b {
			case "X": // lose
				win = 0
				form = 2 // Paper
			case "Y": // draw
				win = 3
				form = 3 // Scissors
			case "Z": // win
				win = 6
				form = 1 // Rock
			}
		}
		totalPoints += win + form
	}
	return totalPoints
}
