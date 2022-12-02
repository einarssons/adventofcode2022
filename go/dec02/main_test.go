package main

import (
	"testing"

	u "github.com/einarssons/adventofcode2022/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	require.Equal(t, 15, task1(lines))
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	require.Equal(t, 12, task2(lines))
}
