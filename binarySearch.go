package main

import "fmt"

func main() {
	slice := []int{1, 5, 15, 20}
	fmt.Println(binarySearch(20, slice))
}

func binarySearch(value int, slice []int) int {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := (low + high) / 2

		if value == slice[mid] {
			return mid
		} else if value > slice[mid] {
			low = mid + 1
		} else if value < slice[mid] {
			high = mid - 1
		}
	}
	return -1
}
