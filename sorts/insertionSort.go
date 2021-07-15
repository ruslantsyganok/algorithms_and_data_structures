package main

import "fmt"

func main() {
	slice := []int{256, 128, 64}
	insertionSort(slice)
	fmt.Println(slice)
}

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		j := i
		for j > 0 {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j = j - 1
		}
	}
}