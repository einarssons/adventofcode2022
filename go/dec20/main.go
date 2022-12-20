package main

import (
	"fmt"

	u "github.com/einarssons/adventofcode2022/go/utils"
)

func main() {
	numbers := u.ReadNumbersFromFile("input")
	fmt.Println("task1: ", task1(numbers))
	fmt.Println("task2: ", task2(numbers))
}

type pair struct {
	val int
	pos int
}

func task1(numbers []int) int {
	pairs := make([]pair, len(numbers))
	for i := 0; i < len(numbers); i++ {
		pairs[i] = pair{numbers[i], i}
	}
	for origPos := 0; origPos < len(numbers); origPos++ {
		currPos := find(origPos, pairs)
		p := pairs[currPos]
		move(pairs, currPos, p.val)
	}
	clear := 0
	zeroPos := findZero(pairs)
	for i := 1000; i <= 3000; i += 1000 {
		pos := (zeroPos + i) % len(numbers)
		p := pairs[pos]
		clear += p.val
	}
	return clear
}

func task2(numbers []int) int {
	key := 811589153
	pairs := make([]pair, len(numbers))
	for i := 0; i < len(numbers); i++ {
		pairs[i] = pair{numbers[i] * key, i}
	}
	for m := 0; m < 10; m++ {
		for origPos := 0; origPos < len(numbers); origPos++ {
			currPos := find(origPos, pairs)
			p := pairs[currPos]
			move(pairs, currPos, p.val)
		}
	}

	clear := 0
	zeroPos := findZero(pairs)
	for i := 1000; i <= 3000; i += 1000 {
		pos := (zeroPos + i) % len(numbers)
		p := pairs[pos]
		clear += p.val
	}
	return clear
}

func find(pos int, pairs []pair) int {
	for i := 0; i < len(pairs); i++ {
		if pairs[i].pos == pos {
			return i
		}
	}
	panic("could not find")
}

func findZero(pairs []pair) int {
	for i := 0; i < len(pairs); i++ {
		if pairs[i].val == 0 {
			return i
		}
	}
	panic("could not find")
}

func newPos(pos, step, length int) int {
	new := pos + step
	l1 := length - 1
	if new >= length {
		nrWraps := new / l1
		new -= nrWraps * l1
		if new == 0 {
			new += l1
		}
	}
	if new < 0 {
		nrWraps := new / l1
		new -= nrWraps * l1
		if new < 0 {
			new += l1
		}
	}
	return new
}

func move(l []pair, pos, step int) {
	val := l[pos]
	newPos := newPos(pos, step, len(l))
	switch {
	case pos == newPos:
		return
	case pos > newPos:
		// leave first newPos-1 element
		// shift [newPos,pos-1] +1
		// insert val at newPos
		// leave elements after pos
		copy(l[newPos+1:pos+1], l[newPos:pos])
		l[newPos] = val
	case pos < newPos:
		// leave first pos-1 element
		// shift [pos+1,newPos+1] -1
		// insert val at newPos
		// leave elements after newPos
		copy(l[pos:newPos], l[pos+1:newPos+1])
		l[newPos] = val
	}
}
