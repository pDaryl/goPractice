package main

import "fmt"

func findMinAndMaxOfSubSlices (slice []int, k int) ([]int, []int) {
	var minValues []int
	var maxValues []int
	
	if len(slice) == 0 || k == 0{
		return minValues, maxValues
	}

	if k > len(slice) {
		min, max := minAndMax(slice)
		maxValues = append(maxValues, max)
		minValues = append(minValues, min)
		return minValues, maxValues
	}

	
	
	for start := 0; start <= len(slice)-k; start++{
		subSlice := slice[start:start + k]
		min, max := minAndMax(subSlice)
		maxValues = append(maxValues, max)
		minValues = append(minValues, min)
	}

	return minValues, maxValues
}


func minAndMax (slice []int) (int, int) {
	if len(slice) == 0 {
		return 0, 0
	}

	min := slice[0]
	max := slice[0]

	for _, num := range slice {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}

	}
	return min, max
}


func main(){
slice := []int {1, 3, 5, 7, 0, 6, 7, 8}
k := 2

minValues, maxValues := findMinAndMaxOfSubSlices(slice, k)

fmt.Println("Minimum values:", minValues)
fmt.Println("Maximum values:", maxValues)


//########################################


slice1 := []int {}
k1 := 2

minValues1, maxValues1 := findMinAndMaxOfSubSlices(slice1, k1)

fmt.Println("Minimum values:", minValues1)
fmt.Println("Maximum values:", maxValues1)

//########################################


slice2 := []int {1, 3, 5, 7, 0, 6, 7, 8}
k2 := 0

minValues2, maxValues2 := findMinAndMaxOfSubSlices(slice2, k2)

fmt.Println("Minimum values:", minValues2)
fmt.Println("Maximum values:", maxValues2)

//########################################


slice3 := []int {1, 3, 5, 7, 0, 6, 7, 8}
k3 := 15

minValues3, maxValues3 := findMinAndMaxOfSubSlices(slice3, k3)

fmt.Println("Minimum values:", minValues3)
fmt.Println("Maximum values:", maxValues3)

}
