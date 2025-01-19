package main

import (
	"fmt"
)

func mergeSlices(slice1 []int, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return slice1
	}

	storeNums := make(map[int] bool)
	var merged []int

	for _, num := range slice1 {
	if !storeNums[num]{
		storeNums[num] = true
		merged = append(merged, num)
	}
	}

	for _, num := range slice2 {
		if !storeNums[num] {
			storeNums[num] = true
			merged = append(merged, num)
		}
	}

return merged
	
}


func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}
  
	fmt.Println(mergeSlices(slice1, slice2))

	//[]int{1, 2, 3, 4, 5, 6}
  
} 

