package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	actual := BubbleSort([]int{5, 4, 3, 2, 1})

	assert.Equal(t, expected, actual)
}
