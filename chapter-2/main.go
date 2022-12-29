package main

import (
	"log"

	"golang.org/x/exp/constraints"
)

func insertionSort[T constraints.Ordered](slice []T) []T {
	sortedSlice := make([]T, len(slice))
	copy(sortedSlice, slice)

	for index, number := range sortedSlice {
		leftIndex := index - 1

		for ; leftIndex >= 0 && sortedSlice[leftIndex] > number; leftIndex-- {
			sortedSlice[leftIndex+1] = sortedSlice[leftIndex]
		}

		sortedSlice[leftIndex+1] = number
	}

	return sortedSlice
}

func insertionSortDecreasing[T constraints.Ordered](slice []T) []T {
	sortedSlice := make([]T, len(slice))
	copy(sortedSlice, slice)

	for index := len(slice) - 1; index >= 0; index-- {
		number := sortedSlice[index]
		rightIndex := index + 1

		for ; rightIndex < len(slice) && sortedSlice[rightIndex] < number; rightIndex++ {
			sortedSlice[rightIndex-1] = sortedSlice[rightIndex]
		}

		sortedSlice[rightIndex-1] = number
	}

	return sortedSlice
}

func insertSortRecursive[T constraints.Ordered](slice []T) []T {
	sortedSlice := make([]T, len(slice))
	copy(sortedSlice, slice)

	lastIndex := len(sortedSlice) - 1

	if lastIndex < 1 {
		return sortedSlice
	}

	sortedSlice = append(insertSortRecursive(sortedSlice[:lastIndex]), sortedSlice[lastIndex])

	index := 0

	for index = 0; sortedSlice[index] < sortedSlice[lastIndex]; index++ {
	}

	lastElements := make([]T, 1)
	lastElements[0] = sortedSlice[lastIndex]
	lastElements = append(lastElements, sortedSlice[index:lastIndex]...)

	sortedSlice = append(sortedSlice[:index], lastElements...)

	return sortedSlice
}

func linearSearch[T comparable](slice []T, find T) int {
	for index, element := range slice {
		if find == element {
			return index
		}
	}

	return -1
}

func binarySearch[T constraints.Ordered](slice []T, find T) int {
	min := 0
	max := len(slice) - 1
	middle := (max + min) / 2

	for max >= min {
		if slice[middle] == find {
			return middle
		}

		if slice[middle] > find {
			max = middle - 1
		} else {
			min = middle + 1
		}
		middle = (max + min) / 2
	}

	return -1
}

type bit bool

const (
	zero bit = false
	one  bit = true
)

func maxMin[T constraints.Ordered](a, b T) (T, T) {
	if a < b {
		return b, a
	}
	return a, b
}

func addBinaryIntegers(a []bit, b []bit) []bit {
	max, min := maxMin(len(a), len(b))
	if max == 0 {
		return []bit{}
	}

	c := make([]bit, max+1)

	if len(a) > len(b) {
		b = append(b, make([]bit, max-min)...)
	} else {
		a = append(a, make([]bit, max-min)...)
	}

	c[0] = a[0] != b[0]
	carry := a[0] && b[0]

	for index := 1; index < max; index++ {
		c[index] = carry != (a[index] != b[index])
		carry = (a[index] && b[index]) || (carry && (a[index] != b[index]))
	}

	c[max] = carry

	return c
}

func selectionSort[T constraints.Ordered](slice []T) []T {
	sortedSlice := make([]T, len(slice))
	copy(sortedSlice, slice)

	for index := range sortedSlice {
		minIndex := index

		for rightIndex := index; rightIndex < len(sortedSlice); rightIndex++ {
			if sortedSlice[rightIndex] < sortedSlice[minIndex] {
				minIndex = rightIndex
			}
		}

		sortedSlice[index], sortedSlice[minIndex] = sortedSlice[minIndex], sortedSlice[index]
	}

	return sortedSlice
}

func merge[T constraints.Ordered](left []T, right []T) []T {
	lenLeft, lenRight := len(left), len(right)

	merged := make([]T, lenLeft+lenRight)

	index := 0
	leftIndex := 0
	rightIndex := 0

	for leftIndex < lenLeft && rightIndex < lenRight {
		if left[leftIndex] < right[rightIndex] {
			merged[index] = left[leftIndex]
			leftIndex++
		} else {
			merged[index] = right[rightIndex]
			rightIndex++
		}

		index++
	}

	for leftIndex < lenLeft {
		merged[index] = left[leftIndex]
		leftIndex++
		index++
	}

	for rightIndex < lenRight {
		merged[index] = right[rightIndex]
		rightIndex++
		index++
	}

	return merged
}

func mergeSort[T constraints.Ordered](slice []T) []T {
	sortedSlice := make([]T, len(slice))
	copy(sortedSlice, slice)

	if len(sortedSlice) <= 1 {
		return sortedSlice
	}

	middle := len(sortedSlice) / 2

	return merge(
		mergeSort(sortedSlice[0:middle]),
		mergeSort(sortedSlice[middle:]),
	)
}

type number interface {
	 constraints.Integer | constraints.Float
}

func sumContains[T number](slice []T, number T) bool {
	if len(slice) <= 1 {
		return false
	}

	index := binarySearch(slice[1:], number-slice[0])
	if index >= 0 {
		return true
	}

	return sumContains(slice[1:], number)
}

func main() {
	array := []int{5, 2, 4, 6, 1, 3}
	log.Println(insertionSort(array))
	log.Println(insertionSortDecreasing(array))

	log.Println(linearSearch(array, 4))
	log.Println(linearSearch(array, 7))

	bits1 := []bit{one, one, one, one, one}
	bits2 := []bit{one, one, one, one, one}

	log.Println(addBinaryIntegers(bits1, bits2))

	log.Println(selectionSort(array))

	array1 := []int{1, 3, 5, 7, 9}
	array2 := []int{2, 4, 6, 8, 10, 12, 14}

	log.Println(merge(array1, array2))
	log.Println(mergeSort(array))

	log.Println(insertSortRecursive(array))

	for _, array := range array2 {
		log.Println(binarySearch(array2, array))
	}
	log.Println(binarySearch(array2, 1))
	log.Println(binarySearch(array2, 15))

  log.Println(sumContains(array1, 0))
  log.Println(sumContains(array1, 1))
  log.Println(sumContains(array1, 9))
  log.Println(sumContains(array1, 11))
  log.Println(sumContains(array2, 6))
  log.Println(sumContains(array2, 10))
  log.Println(sumContains(array2, 26))

	log.Println(array)
}
