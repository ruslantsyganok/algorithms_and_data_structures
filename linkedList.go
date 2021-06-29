package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct{
	head *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) print() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Println(toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

func (l *linkedList) delete(number int) {
	currentPosition := l.head
	for currentPosition.next.data != number {
		currentPosition = currentPosition.next
	}
	currentPosition.next = currentPosition.next.next
	l.length--
}

func main() {
	list := linkedList{}
	node1 := &node{data:48}
	node2 := &node{data:28}
	node3 := &node{data:18}
	node4 := &node{data:77}
	node5 := &node{data:102}
	node6 := &node{data:3}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.prepend(node5)
	list.prepend(node6)
	list.delete(18)
	list.delete(102)
	list.print()
}
