package main

import "fmt"

func  lengthSubSliceTargetSum (slice []int, k int) int {
	if len(slice) == 0 {
		return len(slice)
	}

	currentSum := 0 
	maxLength := 0
	start := 0 

	for end := 0; end < len(slice); end++ {
		currentSum += slice[end]

		for currentSum > k {
			currentSum -= slice[start]
			start++
		}

		if end - start + 1 > maxLength {
			maxLength = end - start + 1
		}
	}
	return maxLength
}


func main(){
slice := []int{1, 1, 1, 2, 2, 2, 3, 4, 5, 6}
k := 6

fmt.Println(lengthSubSliceTargetSum(slice, k))
}
