package main

import "fmt"

func removeDuplicates (slice []int) []int {
	if len(slice) == 0 {
		return []int {}
	}
	var result []int
	mp := make(map[int] bool)

	for _, num := range slice{

		_, ok := mp[num] 
		if !ok {
			result = append(result, num)
			mp[num] = true
		}
	}
	return result
}




func main (){
slice := []int {1, 1, 2, 3, 2, 3, 4, 5, 6, 1, 2, 3, 7}


result := removeDuplicates(slice)

fmt.Println(result)
}