package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}
type LinkedList struct {
	first  *Node
	length int
}

func (list LinkedList) CheckCycles() {
	var head *Node
	head = list.first
	visited := make(map[*Node]bool)
	for n := head; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println("Visited: ", n.Value)
	}
}

func main() {
	ll := &LinkedList{}
	ll.first = &Node{Value: 1}
	ll.first.Next = &Node{Value: 2}
	ll.first.Next.Next = &Node{Value: 3}
	ll.first.Next.Next.Next = &Node{Value: 4}
	//ll.first.Next.Next.Next.Next = ll.first
	n := &Node{Value: 5}
	ll.first.Next.Next.Next.Next = n
	fmt.Println("Cycle exist: ")
	ll.CheckCycles()

}
