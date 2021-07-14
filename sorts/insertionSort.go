package main

import "fmt"

func main() {
	slice := []int{112, 777, 444, 8, 3, 1, 1024, 313}
	insertionSort(slice)
	fmt.Println(slice)
}

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}