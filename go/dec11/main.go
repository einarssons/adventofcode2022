package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("task1: ", task1(initMonkeys()))
	fmt.Println("task2: ", task2(initMonkeys(), 10000))
}

func task1(monkeys []monkey) int {
	for round := 0; round < 20; round++ {
		for m := 0; m < len(monkeys); m++ {
			monkeys[m].run(monkeys)
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
	return monkeys[0].inspections * monkeys[1].inspections
}

func task2(monkeys []monkey, nrRuns int) int {
	for round := 0; round < nrRuns; round++ {
		for m := 0; m < len(monkeys); m++ {
			monkeys[m].run2(monkeys)
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
	return monkeys[0].inspections * monkeys[1].inspections
}

type transform func(int) int

type monkey struct {
	items       []int
	next        [2]int
	tf          transform
	divisor     int
	inspections int
	nr          int
}

func (m *monkey) run(monkeys []monkey) {
	for _, item := range m.items {
		newLevel := m.tf(item)
		newLevel /= 3
		if newLevel%m.divisor == 0 {
			monkeys[m.next[0]].add(newLevel)
		} else {
			monkeys[m.next[1]].add(newLevel)
		}
		m.inspections++
	}
	m.items = m.items[:0]
}

func (m *monkey) run2(monkeys []monkey) {
	primeProduct := 1
	for _, mk := range monkeys {
		primeProduct *= mk.divisor
	}
	for _, item := range m.items {
		newLevel := m.tf(item)
		newLevel = newLevel % primeProduct
		if newLevel%m.divisor == 0 {
			monkeys[m.next[0]].add(newLevel)
		} else {
			monkeys[m.next[1]].add(newLevel)
		}
		m.inspections++
	}
	m.items = m.items[:0]
}

func (m *monkey) add(item int) {
	m.items = append(m.items, item)
}

var testMonkeys = []monkey{
	{
		items:   []int{79, 98},
		tf:      func(a int) int { return a * 19 },
		divisor: 23,
		next:    [2]int{2, 3},
		nr:      0,
	},
	{
		items:   []int{54, 65, 75, 74},
		tf:      func(a int) int { return a + 6 },
		divisor: 19,
		next:    [2]int{2, 0},
		nr:      1,
	},
	{
		items:   []int{79, 60, 97},
		tf:      func(a int) int { return a * a },
		divisor: 13,
		next:    [2]int{1, 3},
		nr:      2,
	},
	{
		items:   []int{74},
		tf:      func(a int) int { return a + 3 },
		divisor: 17,
		next:    [2]int{0, 1},
		nr:      3,
	},
}

func initMonkeys() []monkey {
	return []monkey{
		{
			items:   []int{98, 89, 52},
			tf:      func(a int) int { return a * 2 },
			divisor: 5,
			next:    [2]int{6, 1},
			nr:      0,
		},
		{
			items:   []int{57, 95, 80, 92, 57, 78},
			tf:      func(a int) int { return a * 13 },
			divisor: 2,
			next:    [2]int{2, 6},
			nr:      1,
		},
		{
			items:   []int{82, 74, 97, 75, 51, 92, 83},
			tf:      func(a int) int { return a + 5 },
			divisor: 19,
			next:    [2]int{7, 5},
			nr:      2,
		},
		{
			items:   []int{97, 88, 51, 68, 76},
			tf:      func(a int) int { return a + 6 },
			divisor: 7,
			next:    [2]int{0, 4},
			nr:      3,
		},
		{
			items:   []int{63},
			tf:      func(a int) int { return a + 1 },
			divisor: 17,
			next:    [2]int{0, 1},
			nr:      4,
		},
		{
			items:   []int{94, 91, 51, 63},
			tf:      func(a int) int { return a + 4 },
			divisor: 13,
			next:    [2]int{4, 3},
			nr:      5,
		},
		{
			items:   []int{61, 54, 94, 71, 74, 68, 98, 83},
			tf:      func(a int) int { return a + 2 },
			divisor: 3,
			next:    [2]int{2, 7},
			nr:      6,
		},
		{
			items:   []int{90, 56},
			tf:      func(a int) int { return a * a },
			divisor: 11,
			next:    [2]int{3, 5},
			nr:      7,
		},
	}
}
