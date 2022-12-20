package main

import (
	"fmt"
	"testing"

	u "github.com/einarssons/adventofcode2022/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	numbers := u.ReadNumbersFromFile("test.txt")
	result := task1(numbers)
	require.Equal(t, 3, result)
}

func TestTask2(t *testing.T) {
	numbers := u.ReadNumbersFromFile("test.txt")
	result := task2(numbers)
	require.Equal(t, 1623178306, result)
}

func TestNewPos(t *testing.T) {
	cases := []struct {
		pos    int
		move   int
		length int
		newPos int
	}{
		{6, 6, 7, 6},
		{5, 4, 7, 3},
		{1, -3, 7, 4},
		{1, -9, 7, 4},
		{6, 1, 7, 1},
		{6, 7, 7, 1},
	}

	for i, tc := range cases {
		gotPos := newPos(tc.pos, tc.move, tc.length)
		require.Equal(t, tc.newPos, gotPos, fmt.Sprintf("%d: %v", i, tc))
	}
}
