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
		fmt.Println(head.data)
		head = head.next
		l.length--
	}
}

func (l *linkedList) delete(value int) {
	currentPosition := l.head
	for currentPosition.data != value  {
		currentPosition = currentPosition.next
	}
	currentPosition.next = currentPosition.next.next
	l.length--
}

func main() {
	list := linkedList{}
	n1 := node{data: 48}
	n2 := node{data: 128}
	n3 := node{data: 783}
	n4 := node{data: 1024}
	n5 := node{data: 256}
	list.prepend(&n1)
	list.prepend(&n2)
	list.prepend(&n3)
	list.prepend(&n4)
	list.prepend(&n5)
	list.delete(128)
	list.print()
}
