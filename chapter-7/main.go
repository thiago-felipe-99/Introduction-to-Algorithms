package main

import (
	"log"

	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](slice []T) []T {
	newSlice := make([]T, len(slice))
	copy(newSlice, slice)

	quicksort(newSlice)

	return newSlice
}

func partition[T constraints.Ordered](slice []T) int {
	pivot := slice[len(slice)-1]

	middle := 0

	for index, value := range slice {
		if value < pivot {
			slice[index], slice[middle] = slice[middle], slice[index]
			middle++
		}
	}

	slice[middle], slice[len(slice)-1] = pivot, slice[middle]

	return middle
}

func quicksort[T constraints.Ordered](slice []T) {
	if len(slice) <= 1 {
		return
	}

	middle := partition(slice)
  log.Println(slice)
	quicksort(slice[:middle])
	quicksort(slice[middle+1:])
}

func main() {
	log.Println(QuickSort([]int{2, 8, 7, 1, 3, 5, 6, 4}))
}
