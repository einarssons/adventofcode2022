package main

import (
	"testing"

	u "github.com/einarssons/adventofcode2022/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadRawLinesFromFile("test.txt")
	result := task1(lines)
	require.Equal(t, 6032, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadRawLinesFromFile("test.txt")
	result := task2(lines, 4, true)
	require.Equal(t, 5031, result)
}
