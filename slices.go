package main

import "fmt"

func countFrequencies (slice []int) map[int]int{
	if len(slice) == 0 {
		return make(map[int]int)}

	freqCount := make(map[int] int)

	for _, num := range slice {

		_, ok := freqCount[num]
		if !ok {
		freqCount[num] = 1
		} else {
			freqCount[num]++
		}
	}
	return freqCount
}

func main (){
slice := []int{1, 2, 3, 2, 1, 4, 2}

fmt.Println(countFrequencies(slice))
}