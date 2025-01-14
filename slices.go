package main

import "fmt"

func rotateSlice ( slice []int , k int) []int {

	if len(slice) <= 1 || k == 0 {
		return slice
	}

	steps := k % len(slice)
	if steps == 0 {
		return slice
	}

	part1 := slice[len(slice)-steps:]
	part2 := slice[:len(slice)-steps]

	result := append(part1, part2...)

	return result
}

func main (){
	slice := []int{1, 2, 3, 4, 5}
	k := 2

	fmt.Println(rotateSlice(slice, k))
}