package main

import (
	"fmt"
	"strings"

	"github.com/chrispappas/golang-generics-set/set"
	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("input")
	fmt.Println("task1: ", task1(lines))
	fmt.Println("task2: ", task2(lines))
}

type monkey struct {
	number   int
	a, b, op string
}

func task1(lines []string) int {
	monkeys := make(map[string]monkey)
	monkeysLeft := set.FromSlice([]string{})
	monkeysDone := set.FromSlice([]string{})
	for _, line := range lines {
		k, v := readMonkey(line)
		monkeys[k] = v
		if v.a != "" {
			monkeysLeft.Add(k)
		} else {
			monkeysDone.Add(k)
		}
	}
topLoop:
	for {
		for _, m := range monkeysLeft.Values() {
			mk := monkeys[m]
			if monkeysDone.Has(mk.a) && monkeysDone.Has(mk.b) {
				a := monkeys[mk.a]
				b := monkeys[mk.b]
				var r int
				switch mk.op {
				case "+":
					r = a.number + b.number
				case "-":
					r = a.number - b.number
				case "*":
					r = a.number * b.number
				case "/":
					r = a.number / b.number
				}
				mk.number = r
				monkeys[m] = mk
				monkeysDone.Add(m)
				monkeysLeft.Delete(m)
				if m == "root" {
					break topLoop
				}
			}
		}
	}
	return monkeys["root"].number
}

func task2(lines []string) int {
	monkeys := make(map[string]monkey)
	monkeysLeft := set.FromSlice([]string{})
	monkeysDone := set.FromSlice([]string{})
	for _, line := range lines {
		k, v := readMonkey(line)
		if k == "root" {
			v.op = "-"
			v.number = 0
		}
		monkeys[k] = v
		if k != "humn" && v.a == "" || k == "root" {
			monkeysDone.Add(k)
		} else {
			monkeysLeft.Add(k)
		}
	}
	parents := make(map[string]string)
	for n, m := range monkeys {
		if m.a != "" {
			parents[m.a] = n
		}
		if m.b != "" {
			parents[m.b] = n
		}
	}
phase1:
	for {
		nrLeft := monkeysLeft.Len()
		for _, m := range monkeysLeft.Values() {
			if m == "humn" {
				continue
			}
			mk := monkeys[m]
			if monkeysDone.Has(mk.a) && monkeysDone.Has(mk.b) {
				a := monkeys[mk.a]
				b := monkeys[mk.b]
				var r int
				switch mk.op {
				case "+":
					r = a.number + b.number
				case "-":
					r = a.number - b.number
				case "*":
					r = a.number * b.number
				case "/":
					r = a.number / b.number
				}
				mk.number = r
				monkeys[m] = mk
				monkeysDone.Add(m)
				monkeysLeft.Delete(m)
			}
		}
		nrLeftAfter := monkeysLeft.Len()
		if nrLeftAfter == nrLeft {
			break phase1
		}
	}
phase2:
	for {
		for _, m := range monkeysLeft.Values() {
			mk := monkeys[m]
			p, ok := parents[m]
			if !ok {
				panic(fmt.Sprintf("monkey %s has no parent", m))
			}
			if _, ok := monkeysLeft[p]; ok {
				continue
			}
			pk := monkeys[p]
			isLeft := true
			s := pk.b
			if pk.b == m {
				isLeft = false
				s = pk.a
			}
			if !monkeysDone.Has(s) {
				continue // Other side not ready yet
			}
			if isLeft {
				var a int
				bk := monkeys[s]
				switch pk.op {
				case "+":
					a = pk.number - bk.number
				case "-":
					a = pk.number + bk.number
				case "*":
					a = pk.number / bk.number
				case "/":
					a = pk.number * bk.number
				}
				mk.number = a
			} else {
				var b int
				ak := monkeys[s]
				switch pk.op {
				case "+":
					b = pk.number - ak.number
				case "-":
					b = ak.number - pk.number
				case "*":
					b = pk.number / ak.number
				case "/":
					b = ak.number / pk.number
				}
				mk.number = b
			}
			monkeys[m] = mk
			monkeysDone.Add(m)
			monkeysLeft.Delete(m)
			if m == "humn" {
				break phase2
			}

		}
	}
	return monkeys["humn"].number
}

func readMonkey(line string) (string, monkey) {
	parts := strings.Split(line, " ")
	name := parts[0][:4]
	if len(parts) == 2 {
		return name, monkey{number: u.Atoi(parts[1])}
	}
	return name, monkey{0, parts[1], parts[3], parts[2]}
}
