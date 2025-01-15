package main

import "fmt"

func removeValue ( slice []int , value int) []int {
	if len(slice) == 0 {
		return slice
	}

	var result []int

	for _, num := range slice {
		if num != value {
			result = append(result, num)
		}
	}
return result
}

func main (){
	slice := []int{1, 4, 5, 6, 2}
	value := 2

	fmt.Println(removeValue(slice, value))
}