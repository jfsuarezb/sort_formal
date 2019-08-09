package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testGenerateSlice(t *testing.T) {
	c := require.New(t)
	c.Len(generateSlice(5), 5)
}

func testQuickSort(t *testing.T) {
	c := require.New(t)
	array := []int{9, 7, 8, 5, 6}
	sortedArray := []int{5, 6, 7, 8, 9}
	c.Equal(sortedArray, quicksort(array))
}
