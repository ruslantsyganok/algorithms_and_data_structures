package main

import "fmt"

func main() {
	slice := []int{112, 777, 444, 8, 3, 1, 1024, 313}
	fmt.Println(mergeSort(slice))
}

func mergeSort(slice []int) []int{
	if len(slice) > 1 {
		left := slice[:len(slice)/2]
		right := slice[len(slice)/2:]
		return merge(mergeSort(left), mergeSort(right))
	}
	return slice
}

func merge(left, right []int) []int {
	result := []int{}

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			fmt.Println("Ll:", left[i])
			i++
		} else {
			result = append(result, right[i])
			fmt.Println("Rl:", right[i])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}