package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testRandomArray(t *testing.T) {
	c := require.New(t)
	c.Len(randomArray(), 5)
}

func testSortArray(t *testing.T) {
	c := require.New(t)
	arrayToSort := [5]int{3, 2, 5, 8, 6}
	sortedArray := [5]int{2, 3, 5, 6, 8}
	testArray := sortArray(arrayToSort)
	c.Equal(sortedArray, testArray)
}
