package main

import (
	"testing"

	u "github.com/einarssons/adventofcode2022/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	result := task1(lines, 10)
	require.Equal(t, 26, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	result := task2(lines, 20)
	require.Equal(t, 56000011, result)
}
