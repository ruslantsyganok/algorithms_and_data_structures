package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := []int{777, 444, 8, 1, 112, 1024, 313, 9, 12}
	fmt.Println(quickSort(slice))
}

func quickSort(slice []int) []int{
	if len(slice) > 1 {
		pivotIdx := rand.Intn(len(slice) - 1)
		var left, mid, right []int

		for i := 0; i < len(slice); i++ {
			if slice[i] > slice[pivotIdx] {
				right = append(right, slice[i])
			} else if slice[i] == slice[pivotIdx] {
				mid = append(mid, slice[i])
			} else if slice[i] < slice[pivotIdx] {
				left = append(left, slice[i])
			}
		}
		return append(append(quickSort(left), mid...), quickSort(right)...)
	}
	return slice
}
