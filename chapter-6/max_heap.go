package main

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type MaxHeap[T constraints.Ordered, K any] struct {
	Heap[T, K]
}

func (heap MaxHeap[T, K]) HeapifyRecursive(index int) {
	if heap.Size() <= index {
		return
	}

	left, right := heap.Childrens(index)
	largest := index

	if left < heap.Size() && heap.IsLess(largest, left) {
		largest = left
	}

	if right < heap.Size() && heap.IsLess(largest, right) {
		largest = right
	}

	if largest != index {
		heap.Swap(largest, index)
		heap.HeapifyRecursive(largest)
	}
}

func (heap MaxHeap[T, K]) Heapify(index int) {
	for largest := index; heap.Size() > index; largest = index {
		left, right := heap.Childrens(index)

		if left < heap.Size() && heap.IsLess(largest, left) {
			largest = left
		}

		if right < heap.Size() && heap.IsLess(largest, right) {
			largest = right
		}

		if largest == index {
			break
		}

		heap.Swap(largest, index)
		index = largest
	}
}

func (heap MaxHeap[T, K]) heapifyWithSize(index int, size int) {
	for largest := index; size > index; largest = index {
		left, right := heap.Childrens(index)

		if left < size && heap.IsLess(largest, left) {
			largest = left
		}

		if right < size && heap.IsLess(largest, right) {
			largest = right
		}

		if largest == index {
			break
		}

		heap.Swap(largest, index)
		index = largest
	}
}

func (heap MaxHeap[T, K]) Max() Node[T, any] {
	return heap.Heap[0]
}

func (heap MaxHeap[T, K]) ExtractMax() Node[T, any] {
	max := heap.Max()

	heap.Heap[0] = heap.Heap[heap.Size()-1]
	heap.Heap = heap.Heap[:heap.Size()-1]
	heap.Heapify(0)

	return max
}

func (heap MaxHeap[T, K]) IncreaseKey(index int, newKey T) error {
	if heap.Heap[index].Key > newKey {
		return errors.New("The new key is less than the current key")
	}

	heap.Heap[index].Key = newKey
	parent := heap.Parent(index)

	for index > 0 && heap.IsLess(parent, index) {
		heap.Swap(parent, index)
		index = heap.Parent(index)
	}

	return nil
}

func (heap MaxHeap[T, K]) DeacreaseKey(index int, newKey T) error {
	if heap.Heap[index].Key < newKey {
		return errors.New("The new key is greater than the current key")
	}

	heap.Heap[index].Key = newKey
	heap.Heapify(index)

	return nil
}

func (heap MaxHeap[T, K]) Insert(node Node[T, any]) {
	heap.Heap = append(heap.Heap, node)
	heap.IncreaseKey(heap.Size()-1, node.Key)
}

func (heap MaxHeap[T, K]) Delete(index int) error {
	if index >= heap.Size() || index <= 0 {
		return errors.New("Invalid index")
	}

	heap.Swap(heap.Size()-1, index)
	heap.Heap = heap.Heap[:heap.Size()-1]

	return nil
}

func NewMaxHeapFromSlice[T constraints.Ordered](slice []T) MaxHeap[T, T] {
	maxHeap := MaxHeap[T, T]{NewHeapFromSlice(slice)}

	for index := (len(slice) - 1) / 2; index >= 0; index-- {
		maxHeap.Heapify(index)
	}

	return maxHeap
}
