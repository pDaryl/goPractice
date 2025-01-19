package main

import (
	"fmt"
)

<<<<<<< HEAD
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

=======
func leftOrRightSliceRotation(slice []int, k int, direction string) []int {
	n := len(slice)

	if n == 0 || k == 0{
		return slice
	}

	k = k % n
	if k == 0{
		return slice
	}

	if direction == "right" {
		return append(slice[n - k:], slice[:n - k]...)
	} else if direction == "left" {
		return append(slice[k:], slice[:k]...)
	}
return slice
}


func main(){
slice := []int {1, 2, 3, 4, 5, 6}
k := 2

fmt.Println(leftOrRightSliceRotation(slice, k, "right"))
}
>>>>>>> Function
