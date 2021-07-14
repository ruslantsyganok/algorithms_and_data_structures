package main

import "fmt"

func main() {
	// fib: 0 1 1 2 3 5 8 13
	fmt.Println(fib(6)) // 8
}

func fib (v int) int {
	if v == 0 {
		return 0
	}
	if v == 1 {
		return 1
	}
	return fib(v - 2) + fib(v - 1)
}
