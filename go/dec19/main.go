package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	//fmt.Println("task1: ", task1(lines)) // 1264
	fmt.Println("task2: ", task2(lines[:3]))
}

func task1(lines []string) int {
	totalQuality := 0
	for _, line := range lines {
		bp := parseLine(line)
		fmt.Println(line, bp)
		robots := make([]int, 4)
		robots[ore] = 1
		money := make([]int, 4)
		nrGeodes := work(robots, money, bp.costs, 1, 24)
		fmt.Println("id", bp.id, "nrGeodes", nrGeodes)
		totalQuality += nrGeodes * bp.id
	}
	return totalQuality
}

// The following are tuning parameters to tune the optimization time
// Tuned until the numbers did not increase any longer.
var (
	maxOreRobots          int = 6
	maxClayRobots             = 10
	maxBeforeObsidian         = 18
	nrMinutes                 = 32
	nrRobotsBeforeWaiting     = 3
)

func task2(lines []string) int {
	prod := 1
	fmt.Println("maxOre", maxOreRobots, "maxClay", maxClayRobots,
		"maxBeforeObsidian", maxBeforeObsidian,
		"nrMinutes", nrMinutes, "nrWait", nrRobotsBeforeWaiting)
	for _, line := range lines {
		bp := parseLine(line)
		//fmt.Println(line, bp)
		robots := make([]int, 4)
		robots[ore] = 1
		money := make([]int, 4)
		nrGeodes := work(robots, money, bp.costs, 1, nrMinutes)
		fmt.Println("id", bp.id, "nrGeodes", nrGeodes, bp)
		prod *= nrGeodes
	}
	return prod
}

func work(robots []int, wallet []int, costs [][]int, minute, nrMinutes int) int {
	//fmt.Println(minute, "robots", robots, "wallet", wallet)
	var robPoss []bool
	nrRobots := 0
	for _, r := range robots {
		nrRobots += r
	}
	if minute < nrMinutes && robots[ore] < maxOreRobots && robots[clay] < maxClayRobots {
		robPoss = newRobotPossibilities(wallet, costs, nrRobots)
	}
	// production of money
	for i, nr := range robots {
		wallet[i] += nr
	}
	if minute == nrMinutes {
		return wallet[geode]
	}
	if minute >= maxBeforeObsidian && robots[obsidian] == 0 {
		return 0
	}
	maxGeodes := wallet[geode]
	// try out all the possibilities
	if len(robPoss) > 0 {
		for kind := 0; kind < 4; kind++ {
			if robPoss[kind] {
				newRobots := make([]int, 4)
				copy(newRobots, robots)
				newRobots[kind]++
				newWallet := make([]int, 4)
				copy(newWallet, wallet)
				payFromWallet(costs[kind], newWallet)
				nrGeodes := work(newRobots, newWallet, costs, minute+1, nrMinutes)
				if nrGeodes > maxGeodes {
					//fmt.Println("maxGeodes:", nrGeodes)
					maxGeodes = nrGeodes
				}
			}
		}
	}

	if len(robPoss) == 0 || len(robots) > nrRobotsBeforeWaiting {
		// Try to go down anyway
		nrGeodes := work(robots, wallet, costs, minute+1, nrMinutes)
		if nrGeodes > maxGeodes {
			//fmt.Println("maxGeodes:", nrGeodes)
			maxGeodes = nrGeodes
		}
	}
	return maxGeodes
}

func canAfford(cost []int, money []int) bool {
	for i, c := range cost {
		if money[i] < c {
			return false
		}
	}
	return true
}

func payFromWallet(cost []int, wallet []int) {
	for i, item := range cost {
		wallet[i] -= item
	}
}

func examplePossibilities(money []int, costs [][]int, level int) []bool {
	nextRobot := make([]bool, 4)
	switch level {
	case 1, 2, 3, 5:
		nextRobot[clay] = true
	case 4, 6:
		nextRobot[obsidian] = true
	case 7, 8:
		nextRobot[geode] = true
	default:
		//
	}
	return nextRobot
}

func newRobotPossibilities(money []int, costs [][]int, level int) []bool {
	possible := make([]bool, 4)
	// try to make as many geode first if possible
	if canAfford(costs[geode], money) {
		possible[geode] = true
		return possible
	}
	// Next try to make obsidian
	if canAfford(costs[obsidian], money) {
		possible[obsidian] = true
		return possible
	}

	nonZero := false

	if canAfford(costs[ore], money) {
		nonZero = true
		possible[ore] = true
	}

	if canAfford(costs[clay], money) {
		nonZero = true
		possible[clay] = true
	}

	if nonZero {
		return possible
	}
	return nil
}

type kind int

const (
	ore      kind = 0
	clay     kind = 1
	obsidian kind = 2
	geode    kind = 3
)

func (k kind) String() string {
	switch k {
	case ore:
		return "ore"
	case clay:
		return "clay"
	case obsidian:
		return "obsidian"
	case geode:
		return "geode"
	default:
		panic("bad")
	}
}

type blueprint struct {
	id    int
	costs [][]int
}

func parseLine(line string) blueprint {
	ps := strings.Split(line, ".")
	costs := make([][]int, 0, 4)
	for c := 0; c < 4; c++ {
		costs = append(costs, make([]int, 4))
	}
	pp := strings.Split(ps[0], " ")
	costs[0][0] = u.Atoi(pp[len(pp)-2])
	pp = strings.Split(ps[1], " ")
	costs[1][0] = u.Atoi(pp[len(pp)-2])
	pp = strings.Split(ps[2], " ")
	costs[2][0] = u.Atoi(pp[len(pp)-5])
	costs[2][1] = u.Atoi(pp[len(pp)-2])
	pp = strings.Split(ps[3], " ")
	costs[3][0] = u.Atoi(pp[len(pp)-5])
	costs[3][2] = u.Atoi(pp[len(pp)-2])
	pp = strings.Split(ps[0], " ")
	id := u.Atoi(pp[1][:len(pp[1])-1])
	return blueprint{id: id, costs: costs}
}
