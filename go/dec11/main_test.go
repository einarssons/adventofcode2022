package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask1(t *testing.T) {
	result := task1(testMonkeys)
	require.Equal(t, 10605, result)
}

func TestTask2(t *testing.T) {
	result := task2(testMonkeys, 10000)
	require.Equal(t, 2713310158, result)
}
