package main

import "fmt"

type linkedList struct {
	head   *node
	length int
}

type node struct {
	data int
	next *node
}

func (l *linkedList) prepend(node *node) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length++
}

func (l linkedList) print() {
	head := l.head
	for l.length != 0 {
		fmt.Printf("%v", head.data)
		fmt.Printf(" ")
		head = head.next
		l.length--
	}
}

func (l *linkedList) delete(value int) {
	// fixes error if we want to delete an element from an empty list
	if l.length == 0 {
		return
	}

	// fixes error when we try to delete a head
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	currentPosition := l.head
	for currentPosition.next.data != value  {
		currentPosition = currentPosition.next
		// fixes error if we want to delete a number that doesn't exist in a list
		if currentPosition.next == nil {
			return
		}
	}

	currentPosition.next = currentPosition.next.next
	l.length--
}

func main() {
	list := linkedList{}
	n1 := node{data: 48}
	n2 := node{data: 128}
	n3 := node{data: 666}
	n4 := node{data: 777}

	list.prepend(&n1)
	list.prepend(&n2)
	list.prepend(&n3)
	list.prepend(&n4)
	list.delete(48)
	list.delete(48)
	list.delete(772453345)
	list.print()

}
