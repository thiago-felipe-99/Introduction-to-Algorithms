package main

import (
	"log"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered, K any] struct {
	Key   T
	Value K
}

func (node Node[T, K]) IsGreater(compare Node[T, K]) bool {
	return node.Key > compare.Key
}

func (node Node[T, K]) IsLess(compare Node[T, K]) bool {
	return node.Key < compare.Key
}

type Heap[T constraints.Ordered, K any] []Node[T, any]

func (heap Heap[T, K]) Left(index int) int {
	return index*2 + 1
}

func (heap Heap[T, K]) Right(index int) int {
	return index*2 + 2
}

func (heap Heap[T, K]) Childrens(index int) (int, int) {
	return heap.Left(index), heap.Right(index)
}

func (heap Heap[T, K]) Parent(index int) int {
	return (index - 1) / 2
}

func (heap Heap[T, K]) Size() int {
	return len(heap)
}

func (heap Heap[T, K]) Keys() []T {
	keys := make([]T, len(heap))

	for index, node := range heap {
		keys[index] = node.Key
	}

	return keys
}

func (heap Heap[T, K]) Values() []K {
	values := make([]K, len(heap))

	for index, node := range heap {
		values[index] = node.Value.(K)
	}

	return values
}

func (heap Heap[T, K]) IsMax() bool {
	return heap.isMax(0)
}

func (heap Heap[T, K]) isMax(index int) bool {
	size := heap.Size()

	if index >= size {
		return true
	}

	left, right := heap.Childrens(index)

	if (left < size && heap[index].IsLess(heap[left])) ||
		(right < size && heap[index].IsLess(heap[left])) {
		return false
	}

	return heap.isMax(left) && heap.isMax(right)
}

func (heap Heap[T, K]) IsMin() bool {
	return heap.isMin(0)
}

func (heap Heap[T, K]) isMin(index int) bool {
	size := heap.Size()

	if index >= size {
		return true
	}

	left, right := heap.Childrens(index)

	if (left < size && heap[index].IsGreater(heap[left])) ||
		(right < size && heap[index].IsGreater(heap[right])) {
		return false
	}

	return heap.isMin(left) && heap.isMin(right)
}

func (heap Heap[T, K]) Swap(index1, index2 int) {
	heap[index1], heap[index2] = heap[index2], heap[index1]
}

func (node Heap[T, K]) IsGreater(index1, index2 int) bool {
	return node[index1].IsGreater(node[index2])
}

func (node Heap[T, K]) IsLess(index1, index2 int) bool {
	return node[index1].IsLess(node[index2])
}

func NewHeapFromSlice[T constraints.Ordered](slice []T) Heap[T, T] {
	heap := make(Heap[T, T], len(slice))

	for index, value := range slice {
		heap[index] = Node[T, any]{value, value}
	}

	return heap
}

func HeapSort[T constraints.Ordered](slice []T) []T {
	heap := NewMaxHeapFromSlice(slice)

	for index := heap.Size() - 1; index >= 0; index-- {
		heap.Swap(index, 0)
		heap.heapifyWithSize(0, index)
    log.Println(heap.Keys())
    log.Println(heap.Values())
	}

	return heap.Keys()
}
