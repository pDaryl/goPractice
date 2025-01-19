package main

import "fmt"

func removeDuplicateElements(slice []int) []int {
	if len(slice) == 0 {
		return slice 
	}

	var uniqueSL []int
	uniqueNum := make(map[int] bool)

	for _, num := range slice {
		if !uniqueNum[num] {
			uniqueNum[num] = true
			uniqueSL = append(uniqueSL, num)
		}
	}
	return uniqueSL
}


func main(){
	slice1 := []int {1, 1, 1, 1, 1}
	slice2 := []int {}
	slice3 := []int {1, 1, 1, 2, 2, 3, 4, 2, 3, 3, 4, 4}

	fmt.Println(removeDuplicateElements(slice1))
	fmt.Println(removeDuplicateElements(slice2))
	fmt.Println(removeDuplicateElements(slice3))
	
}
