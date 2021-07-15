package main

import "fmt"

type Set struct {
	set map[int]bool
}

func (s *Set) Insert(value int) {
	if s.set[value] {
		fmt.Println("The element is already in the set")
		return
	}
	s.set[value] = true
}

func (s *Set) Delete(value int) {
	if s.set[value] {
		delete(s.set, value)
		return
	}
	fmt.Println("There is no such an element")
}

func (s *Set) Size() int {
	return len(s.set)
}

func (s *Set) Show() {
	for k, v := range s.set {
		fmt.Println(k, v)
	}
}

func main() {
	set := Set{set: map[int]bool{}}
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)
	set.Insert(3)
	set.Delete(3)
	fmt.Println(set.Size())
	set.Show()
}
