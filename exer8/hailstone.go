package exer8

/**
* Exercise 8 on CMPT 383. Introduction to Go programming language.
* Append is longer when it is compared to allocate, which is expected.
* It happens because the size in append is not fixed while allocated has a fixed size array
* Append takes 3130 ns/op when Allocate only takes 2590 ns/op
* The difference is not significant in the example, but it is significant if you run the function with multiple runs.
* Therefore, to allocate the array size first is better rather than growing the array.
*/

func Hailstone(hailNum uint) uint {
    if hailNum % 2 == 0 {
		return hailNum / 2
	} else {
		return  3 * hailNum + 1
	}
}

func HailstoneSequenceAppend(num uint) []uint{
	var arr []uint
	for num != 1 {
	   arr = append(arr,num)
	   num = Hailstone(num)
	}
	arr = append(arr,1)
	return arr
}

func countHailLength (num uint) uint{
	// base case
	if num == 1{
		return 1
	} else { 
		return 1 + countHailLength(Hailstone(num))
	}
}

func HailstoneSequenceAllocate(num uint) []uint{
	arr := make([]uint, countHailLength(num))
	for ii := 0; ii < (len(arr) - 1); ii++ {
		arr[ii] = num
		num = Hailstone(num)
	 }
	 arr[len(arr) - 1] = 1
	 return arr
}