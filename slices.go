package main

import (
	"fmt"
	"sort"
)

func mergingInterval(slice [][]int ) [][]int {
	if len(slice) == 0 {
		return slice
	}


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

