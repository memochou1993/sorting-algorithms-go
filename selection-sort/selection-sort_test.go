package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	actual := SelectionSort([]int{5, 4, 3, 2, 1})

	assert.Equal(t, expected, actual)
}
