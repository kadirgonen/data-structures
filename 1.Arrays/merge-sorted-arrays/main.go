package main

import (
	"fmt"
)

func main() {

	sl1 := []int{1, 3, 5, 9, 21, 34}
	sl2 := []int{0, 4, 5, 6}

	fmt.Println(mergeSortedSlices(sl1, sl2))
}

func mergeSortedSlices(slice1, slice2 []int) []int {
	result := make([]int, 0, len(slice1)+len(slice2))

	i, j := 0, 0

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] > slice2[j] {
			result = append(result, slice2[j])
			j++
		} else {
			result = append(result, slice1[i])
			i++
		}
	}

	for ; i < len(slice1); i++ {
		result = append(result, slice1[i])
	}
	for ; j < len(slice2); j++ {
		result = append(result, slice2[j])
	}

	return result
}