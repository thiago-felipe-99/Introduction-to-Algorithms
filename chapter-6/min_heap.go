package main

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered, K any] struct {
	Heap[T, K]
}

func (heap MinHeap[T, K]) Heapify(index int) {
	if heap.Size() <= index {
		return
	}

	left, right := heap.Childrens(index)
	smaller := index

	if left < heap.Size() && heap.IsGreater(smaller, left) {
		smaller = left
	}

	if right < heap.Size() && heap.IsGreater(smaller, right) {
		smaller = right
	}

	if smaller != index {
		heap.Swap(smaller, index)
		heap.Heapify(smaller)
	}
}

func (heap MinHeap[T, K]) HeapifyIterative(index int) {
	for smaller := index; heap.Size() > index; smaller = index {
		left, right := heap.Childrens(index)

		if left < heap.Size() && heap.IsGreater(smaller, left) {
			smaller = left
		}

		if right < heap.Size() && heap.IsGreater(smaller, right) {
			smaller = right
		}

		if smaller == index {
			break
		}

		heap.Swap(smaller, index)
		index = smaller
	}
}

func (heap MinHeap[T, K]) heapifyWithSize(index int, size int) {
	for smaller := index; size > index; smaller = index {
		left, right := heap.Childrens(index)

		if left < size && heap.IsGreater(smaller, left) {
			smaller = left
		}

		if right < size && heap.IsGreater(smaller, right) {
			smaller = right
		}

		if smaller == index {
			break
		}

		heap.Swap(smaller, index)
		index = smaller
	}
}

func (heap MinHeap[T, K]) Min() Node[T, any] {
	return heap.Heap[0]
}

func (heap MinHeap[T, K]) ExtractMin() Node[T, any] {
	min := heap.Min()

	heap.Heap[0] = heap.Heap[heap.Size()-1]
	heap.Heap = heap.Heap[:heap.Size()-1]
	heap.HeapifyIterative(0)

	return min
}

func (heap MinHeap[T, K]) DecreaseKey(index int, newKey T) error {
	if heap.Heap[index].Key > newKey {
		return errors.New("The new key is greater than the current key")
	}

	heap.Heap[index].Key = newKey
	parent := heap.Parent(index)

	for index > 0 && heap.IsGreater(parent, index) {
		heap.Swap(parent, index)
		index = heap.Parent(index)
	}

	return nil
}

func (heap MinHeap[T, K]) IncreaseKey(index int, newKey T) error {
	if heap.Heap[index].Key > newKey {
		return errors.New("The new key is less than the current key")
	}

	heap.Heap[index].Key = newKey
	heap.HeapifyIterative(index)

	return nil
}

func (heap MinHeap[T, K]) Insert(node Node[T, any]) {
	heap.Heap = append(heap.Heap, node)
	heap.DecreaseKey(heap.Size()-1, node.Key)
}

func NewMinHeapFromSlice[T constraints.Ordered](slice []T) MinHeap[T, T] {
	minHeap := MinHeap[T, T]{NewHeapFromSlice(slice)}

	for index := (len(slice) - 1) / 2; index >= 0; index-- {
		minHeap.Heapify(index)
	}

	return minHeap
}
