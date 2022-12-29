package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeapify(t *testing.T) {
	heap := NewMaxHeapFromSlice([]int {16, 4, 10, 14, 7, 9, 3, 2, 8, 1})
	equal := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

	heap.Heapify(1)

	assert.Equal(t, heap.Values(), equal)
}

func TestIsMax(t *testing.T) {
	heap1 := NewHeapFromSlice([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})
	assert.True(t, heap1.IsMax())

	heap2 := NewHeapFromSlice([]int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1})
	assert.False(t, heap2.IsMax())
}

func TestIsMin(t *testing.T) {
	heap1 := NewHeapFromSlice([]int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16})
	assert.True(t, heap1.IsMin())

	heap2 := NewHeapFromSlice([]int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1})
	assert.False(t, heap2.IsMin())
}

func TestNewMaxHeapFromSlice(t *testing.T) {
	slice := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap := NewMaxHeapFromSlice(slice)

	equal := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

	assert.Equal(t, heap.Values(), equal)
	assert.NotEqual(t, slice, equal)
}

func TestNewMinHeapFromSlice(t *testing.T) {
	slice := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap := NewMinHeapFromSlice(slice)

	equal := []int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16}

	assert.Equal(t, heap.Values(), equal)
	assert.NotEqual(t, slice, equal)
}

func TestHeapSort(t *testing.T) {
	slice := []int{1, 4, 5, 10, 6, 7, 9, 8, 2, 3}
	equal := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, HeapSort(slice), equal)
}
