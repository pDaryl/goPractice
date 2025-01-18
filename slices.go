package main

import (
	"fmt"
	"sort"
)

<<<<<<< HEAD
func mergingInterval(slice [][]int ) [][]int {
=======
func zeroToEnd(slice []int) []int {
>>>>>>> Function
	if len(slice) == 0 {
		return slice 
	}

<<<<<<< HEAD

	sort.Slice(slice, func(a int, b int) bool {
		return slice[a][0] < slice[b][0]
	})

	var merged [][]int

	for _, interval := range slice {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0]{
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}

	return merged
}


func main() {
	slice := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	fmt.Println(mergingInterval(slice))
}

=======
	write := 0

	for read := 0; read < len(slice); read++ {
		if slice[read] != 0 {
			slice[write] = slice[read]
			write++
		}else{
			slice[read]++
		}
	}

	for i := write; i < len(slice); i++ {
		slice[i] = 0
	}
	return slice
}

func main (){
slice := []int {1, 0, 2, 0, 3, 0, 4}

fmt.Println(zeroToEnd(slice))
}
>>>>>>> Function
