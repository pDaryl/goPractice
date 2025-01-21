package main

import "fmt"


func reverse(slice []int, start int, end int) {
	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}
}

func rotateSliceInPlace(slice []int, k int) []int {
	len := len(slice)

	if len == 0 || k == 0 {
		return slice
	}

	k = k % len
	
	if k == 0 {
		return slice
	}

	reverse(slice, 0, len - 1) // reverse entire slice
	reverse(slice, 0, k-1) // get the first elements
	reverse(slice, k, len-1) // get the remaining elements

	return slice
}

func main(){
slice := []int {1, 2, 3, 4, 5}
k := 2
fmt.Println(rotateSliceInPlace(slice, k))

// -------------------------------

slice1 := []int {1, 2, 3, 4, 5}
k1 := 0
fmt.Println(rotateSliceInPlace(slice1, k1))

// -----------------------------

slice2 := []int {}
k2 := 2
fmt.Println(rotateSliceInPlace(slice2, k2))

// ----------------------------

slice3 := []int {1, 2, 3, 4, 5}
k3 := 7
fmt.Println(rotateSliceInPlace(slice3, k3))

}
