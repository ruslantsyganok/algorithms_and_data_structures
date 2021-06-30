package main

import "fmt"

type stack struct {
	items []int
}

func (s *stack) push(value int) {
	s.items = append(s.items, value)
}

func (s *stack) pop() int {
	if len(s.items) == 0 {
		return -1
	}
	l := len(s.items) - 1
	removedValue := s.items[l]
	s.items = s.items[:l]
	return removedValue
}

func main() {
	s := stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	fmt.Println(s)
}
