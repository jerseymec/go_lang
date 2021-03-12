package main

import "fmt"

var count int = 0

//Node

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search
func (n *Node) Search(k int) bool {
	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(78)
	tree.Insert(178)
	tree.Insert(8)
	tree.Insert(780)

	fmt.Println(tree)
	fmt.Println(tree.Search(178))
	fmt.Println(count)
}
