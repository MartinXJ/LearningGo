package exer9

import (
	"math/rand"
	"time"
)

type PartialSum struct {
   pMean int
   pStdDev int
}

func RandomArray(length int, maxInt int) []int {

	arr :=  make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for ii := 0; ii < (length - 1); ii++ {
		arr[ii] = (rand.Intn(maxInt))
	}
	return arr

}

/*
func calcMeanStdDev (c chan PartialSum, arr []int){

	mean := 0
	stddev := 0
	for i := 0; i < len(arr); i++ {
		mean = i
		stddev = i
	}
}
*/

// Unfinished
// Not enough time

func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size!")
	}
	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.

	// Initialize two results that will return from MeanStdDev function
	meanTotal := 0.0
	stdDevTotal := 0.0

	// Create a channel for PartialSum
	// c := make (chan PartialSum)

	// Divide the arrays into evenly-sized slices
	// arrLength    := len(arr)
	/*
	var chunksLength int = arrLength / chunks
	arrayChunks := 0
	
	// chunks is a total number of chunks
	for i := 1; i <= chunks; i++ {
        tempArr := [chunksLength]int
		for j := 0; j < chunksLength; j++ {
			tempArr[j] = arr[i * arrayChunks + j]
		}
		go calcMeanStdDev(c, tempArr)
		arrayChunks++ 
	}

*/
	return meanTotal, stdDevTotal
}
