package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func (n *Node) insert(v int) {
	if v > n.Data {
		if n.Right == nil {
			n.Right = &Node{Data: v}
		} else {
			n.Right.insert(v)
		}
	} else if v < n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: v}
		} else {
			n.Left.insert(v)
		}
	}
}

func (n *Node) search(v int) bool {
	if n == nil {
		return false
	}

	if v > n.Data {
		return n.Right.search(v)
	}

	if v < n.Data {
		return n.Left.search(v)
	}
	return true
}

func (n *Node) delete(v int) {

}

func main() {
	tree := &Node{Data: 100}
	fmt.Println(tree)
	tree.insert(125)
	tree.insert(150)
	tree.insert(310)
	fmt.Println(tree.search(310))
}
