package main

import "fmt"

func findSumContinuousSlices (slice []int, k int) [][]int {
	if len(slice) == 0 {
		return [][]int {}
	}

	var result = [][]int {}
	var currentSum int 

	start := 0 // marks the start of a possible sub-array

	for reader := 0; reader < len(slice); reader++{
		currentSum += slice[reader]

		for currentSum > k && start <= reader {
			currentSum -= slice[start]
			start++
		} 

		if currentSum == k {
			subSlice := slice[start:reader+1]
			result = append(result, subSlice)
		}
	}
return result
}

func main(){
slice := []int{1, 2, 3, 4, 5}
k := 6

fmt.Println(findSumContinuousSlices(slice, k))
}
