package main

import "fmt"

func moveZeroToEnd(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	write := 0

	for read := 0; read < len(slice); read++ {
		if slice[read] != 0 {
			slice[write] = slice[read]
			write++
		}
	}

	for i := write; i < len(slice); i++ {
		slice[i] = 0
	}
return slice 
}

func main (){
	slice := []int {1, 2, 0, 3, 0, 4}

	fmt.Println(moveZeroToEnd(slice))
}