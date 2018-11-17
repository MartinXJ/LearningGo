package exer9

import (
	"math/rand"
	// "sort"
)

// Partition the slice arr around a random pivot (in-place), and return the pivot location.
func partition(arr []float64) int {
	// Adapted from https://stackoverflow.com/a/15803401/6871666
	left := 0
	right := len(arr) - 1

	// Choose random pivot
	pivotIndex := rand.Intn(len(arr))

	// Stash pivot at the right of the slice
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Move elements smaller than the pivot to the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last-smaller element
	arr[left], arr[right] = arr[right], arr[left]
	return left
}

func swap(value1 float64, value2 float64) (float64,float64) {
	tempValue := value1
	value1 = value2
	value2 = tempValue
	return value1, value2

}

func InsertionSort(arr []float64) {
	// TODO: implement insertion sort
	if len(arr) >= 2 {
		for ii := 0; ii < len(arr); ii++ {
			index := ii
			for (index != 0 && arr[index] < arr[index-1]){
				arr[index], arr[index-1] = swap(arr[index], arr[index-1])
				index = index - 1
			}
		}
	}
}

const insertionSortCutoff = 20

func QuickSort(arr []float64) {
	QuickSortHelper(arr, 0, len(arr) - 1)
}

func QuickSortHelper(arr [] float64, first int, last int){
	if len(arr) > insertionSortCutoff && (first - last) >= 2 {
		pivotInd := partition(arr)
		// QuickSortHelper(arr, first, pivotInd - 1)
		QuickSortHelper(arr, pivotInd + 1, last)
	} else if len(arr) >= 2 {
		InsertionSort(arr)
	}
}
