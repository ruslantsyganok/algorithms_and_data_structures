package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(i int) {
	h.array = append(h.array, i)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}

	extracted := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]


	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	if len(h.array) == 0 {
		return
	}

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)

	comparisonValue := 0

	for l <= lastIndex {
		if l == lastIndex {
			comparisonValue = l
		} else if h.array[l] > h.array[r] {
			comparisonValue = l
		} else {
			comparisonValue = r
		}

		if h.array[index] < h.array[comparisonValue] {
			h.swap(index, comparisonValue)
			index = comparisonValue
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[index] > h.array[parent(index)] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	heap := &MaxHeap{array: []int{}}
	slice := []int{10, 20}
	for _, v := range slice {
		heap.Insert(v)
		fmt.Println(heap)
	}

	heap.Extract()
	fmt.Println(heap)
}
