package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	line := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	result := findPos(line, 4)
	require.Equal(t, 5, result)
}
