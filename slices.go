package main

import "fmt"

func lengthLongestSubSliceOfSum (slice []int, k int) (int) {
	if len(slice) == 0 {
		return 0
	}

	maxLength := 0
	currentSum := 0
	start := 0

	for end := 0; end < len(slice); end++{
		currentSum += slice[end]
		
		if currentSum > k && start <= end{		
			currentSum -= slice[start]
			start++
		}
		
		if len(slice[start:end+1]) > maxLength{
			maxLength = len(slice[start:end+1])
		}
	}
		return maxLength
}

func main(){
slice := []int{2, 1, 5, 2, 3, 2}
k := 7

fmt.Println(lengthLongestSubSliceOfSum(slice, k))
}
