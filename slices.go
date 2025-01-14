package main

import "fmt"

func findIntersectingElements (slice1 []int, slice2 []int) []int {
	if len(slice1) == 0 || len(slice2) == 0 {
		return []int {}
	}

	var result []int
	freqCount := make(map[int] int)

	for _, num := range slice1 {
		freqCount[num]++
	}
		
	for _, num := range slice2 {
			if freqCount[num] > 0 {
				result = append(result, num)
				freqCount[num]--
			}
	}

return result
}

func main (){
	slice1 := []int{1, 2, 2, 3, 4}
	slice2 := []int{2, 2, 3, 5}

	fmt.Println(findIntersectingElements(slice1, slice2));
}