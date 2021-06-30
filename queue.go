package main

import "fmt"

type queue struct {
	items []int
}

func (q *queue) enqueue (value int) {
	q.items = append(q.items, value)
}

func (q *queue) dequeue () int {
	if len(q.items) == 0 {
		return -1
	}
	removedValue := q.items[0]
	q.items = q.items[1:]
	return removedValue
}

func main() {
	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	fmt.Println(q)
}
