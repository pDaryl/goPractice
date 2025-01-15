package main

import "fmt"

func SliceOfUniqueElements (slice1 []int, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return []int {}
	}
	
	var uniqueElements []int
	mp := make(map[int] bool)

	for _, num := range slice1 {
		mp[num] = true
	}

	for _, num := range slice2 {
		mp[num] = true
	}

	for key := range mp {
		uniqueElements = append(uniqueElements, key)
	}

	return uniqueElements
}

func main (){
	slice1 := []int{1, 2, 3, 4}
slice2 := []int{3, 4, 5, 6}

// Expected output: [1, 2, 3, 4, 5, 6]


	fmt.Println(SliceOfUniqueElements(slice1, slice2))
}