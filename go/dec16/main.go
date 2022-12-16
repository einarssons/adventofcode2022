package main

import (
	"fmt"
	"strings"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

func task1(lines []string) int {
	valves := make(map[string]valve)
	for _, line := range lines {
		valve := parseLine(line)
		valves[valve.name] = valve
	}
	nrMinutesLeft := 30
	var nonZeroValves []string
	for name, v := range valves {
		if v.flow != 0 {
			nonZeroValves = append(nonZeroValves, name)
		}
	}
	allDistances := findAllDistances(valves)
	totalFlow := openValves("AA", nrMinutesLeft, nonZeroValves, valves, allDistances)
	return totalFlow
}

func task2(lines []string) int {
	valves := make(map[string]valve)
	for _, line := range lines {
		valve := parseLine(line)
		valves[valve.name] = valve
	}
	nrMinutesLeft := 26
	var nonZeroValves []string
	for name, v := range valves {
		if v.flow != 0 {
			nonZeroValves = append(nonZeroValves, name)
		}
	}
	bestTotal := 0
	// Split into two groups and let the elefent take one and me the other.
	// Half of the partitions is enough, since symmetric between
	// me and elephant.
	nrPartitions := (1 << len(nonZeroValves)) / 2
	fmt.Printf("task2: will iterate %d times\n", nrPartitions)
	for i := 0; i < nrPartitions; i++ {
		var myZeroValves []string
		var elZeroValves []string
		for n := 0; n < len(nonZeroValves); n++ {
			if (1<<n)&i != 0 {
				myZeroValves = append(myZeroValves, nonZeroValves[n])
			} else {
				elZeroValves = append(elZeroValves, nonZeroValves[n])
			}
		}
		// Visiting at most 3 nodes per characters it too little
		if len(myZeroValves) <= 3 || len(elZeroValves) <= 3 {
			continue
		}
		allDistances := findAllDistances(valves)
		myFlow := openValves("AA", nrMinutesLeft, myZeroValves, valves, allDistances)
		elFlow := openValves("AA", nrMinutesLeft, elZeroValves, valves, allDistances)
		total := myFlow + elFlow
		if total > bestTotal {
			bestTotal = total
		}
		if i%1000 == 0 && i > 0 {
			fmt.Printf("  * iteration %d lowest: %d\n", i, bestTotal)
		}
	}
	return bestTotal
}

func openValves(pos string, minutesLeft int, nonZeroValves []string, valves map[string]valve,
	allDistances map[string]map[string]int) (flow int) {
	maxFlow := 0
	for i, nz := range nonZeroValves {
		minLeft := minutesLeft - allDistances[pos][nz] // walk
		minLeft--                                      // open
		if minLeft <= 0 {
			continue
		}
		flow := valves[nz].flow * minLeft
		if len(nonZeroValves) == 1 {
			if flow > maxFlow {
				maxFlow = flow
			}
			break
		}
		nonZeroLefts := make([]string, 0, len(nonZeroValves)-1)
		for j := 0; j < i; j++ {
			nonZeroLefts = append(nonZeroLefts, nonZeroValves[j])
		}
		for j := i + 1; j < len(nonZeroValves); j++ {
			nonZeroLefts = append(nonZeroLefts, nonZeroValves[j])
		}
		flow += openValves(nz, minLeft, nonZeroLefts, valves, allDistances)
		if flow > maxFlow {
			maxFlow = flow
		}
	}
	return maxFlow
}

func findAllDistances(valves map[string]valve) map[string]map[string]int {
	m := make(map[string]map[string]int)
	for k := range valves {
		m[k] = calcPathCosts(k, valves)
	}
	return m
}

type pathCost struct {
	cost    int
	visited bool
	done    bool
}

func calcPathCosts(start string, valves map[string]valve) map[string]int {
	m := make(map[string]int)
	cm := make(map[string]pathCost)
	for k := range valves {
		cm[k] = pathCost{}
	}
	cm[start] = pathCost{0, true, false}

	for {
		lowestNode, cost, ok := findLowestNonVisited(cm, valves)
		if !ok {
			break
		}
		cm[lowestNode] = pathCost{cost, true, false}
		for k, v := range cm {
			if v.visited && !v.done {
				node := valves[k]
				done := true
				for _, neighborName := range node.neighbors {
					if !cm[neighborName].visited {
						done = false
						break
					}
				}
				if done {
					cm[k] = pathCost{cm[k].cost, true, true}
				}
			}
		}
	}

	for k, v := range cm {
		m[k] = v.cost
	}
	return m
}

func findLowestNonVisited(cm map[string]pathCost, valves map[string]valve) (string, int, bool) {
	lowestCost := u.MaxInt
	lowestNode := ""

	for k, v := range cm {
		if v.visited && !v.done {
			for _, n := range valves[k].neighbors {
				if !cm[n].visited {
					if cm[k].cost+1 < lowestCost {
						lowestNode = n
						lowestCost = cm[k].cost + 1
					}
				}
			}
		}
	}
	ok := lowestNode != ""
	return lowestNode, lowestCost, ok
}

type valve struct {
	name      string
	flow      int
	neighbors []string
}

func parseLine(line string) valve {
	parts := strings.SplitN(line, " ", 10)
	neighbors := strings.Split(parts[9], ",")
	for i := range neighbors {
		neighbors[i] = strings.TrimSpace(neighbors[i])
	}
	flow := u.Atoi(parts[4][5 : len(parts[4])-1])
	return valve{parts[1], flow, neighbors}
}
