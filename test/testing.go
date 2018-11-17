package main

import (
  "fmt"
  "math/rand"
)

func main() {
    arr := []float64{5, 3, 2, 1, 4, 7, 8 , 9 , 10, 12, 12, 11, 10, 9}
	fmt.Println(arr)
	//InsertionSort(arr)
	QuickSort(arr)
	fmt.Println(arr)

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

const insertionSortCutoff = 50

func QuickSort(arr []float64) {
	QuickSortHelper(arr, 0, len(arr) - 1)
}

func QuickSortHelper(arr [] float64, first int, last int){
	if len(arr) > insertionSortCutoff {
		pivotInd := partition(arr)
		fmt.Println(pivotInd)
		// QuickSortHelper(arr, first, pivotInd - 1)
		QuickSortHelper(arr, pivotInd + 1, last)
	} else if len(arr) >= 2 {
		InsertionSort(arr)
	} else {
		return
	}
}
