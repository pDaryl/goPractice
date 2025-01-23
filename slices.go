package main

import "fmt"

func  findLengthLongestSubSliceKElem(slice []int, k int) int {
if len(slice) == 0 || k > len(slice) {
	return len(slice)
}

maxLength := 0
distinctKeys := make(map[int] int)
start := 0

for end := 0; end < len(slice); end++ {
	_, ok := distinctKeys[slice[end]]
	if !ok {
		distinctKeys[slice[end]] = 1
	}else {
		distinctKeys[slice[end]]++
	}

	for len(distinctKeys) > k {
		distinctKeys[slice[start]]--
		if distinctKeys[slice[start]] == 0 {
			delete(distinctKeys, slice[start])
		}
		start++
	}
	if end - start + 1 > maxLength {
		maxLength = end - start + 1
	}
}
return maxLength
}


func main(){
slice := []int{1, 1, 1, 2, 2, 2, 3, 4, 5, 6}
k := 3

fmt.Println(findLengthLongestSubSliceKElem(slice, k))
}
