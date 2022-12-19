package main

import (
	"testing"

	u "github.com/einarssons/adventofcode2022/go/utils"
	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	result := task1(lines)
	require.Equal(t, 33, result)
}

func TestTask2(t *testing.T) {
	lines := u.ReadLinesFromFile("test.txt")
	result := task2(lines)
	// These are not the optimal parameters
	maxClayRobots = 7
	maxOreRobots = 5
	maxBeforeObsidian = 10
	nrRobotsBeforeWaiting = 2
	require.Equal(t, 56*62, result)
}
