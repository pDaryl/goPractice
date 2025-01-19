package main

import "fmt"

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