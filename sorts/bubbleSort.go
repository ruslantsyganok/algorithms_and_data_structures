package main

import "fmt"

func main() {
	slice := []int{112, 777, 444, 8, 3, 1, 1024, 313}
	bubbleSort(slice)
	fmt.Println(slice)
}

func bubbleSort(slice []int) {
	for i := 0; i < len(slice) - 1; i++ {
		for j := 0; j < len(slice) - 1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
