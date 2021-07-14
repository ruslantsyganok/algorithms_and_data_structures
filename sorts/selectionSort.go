package main

import "fmt"

func main() {
	slice := []int{112, 777, 444, 8, 3, 1, 1024, 313}
	selectionSort(slice)
	fmt.Println(slice)
}

func selectionSort(slice []int){
	for i := 0; i < len(slice); i++ {
		var minIdx = i
		for j := i; j < len(slice); j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
}