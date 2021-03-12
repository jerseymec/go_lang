package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func main() {
	header := &LinkedList{Value: 1}
	// header.Value := 1
	tail := &LinkedList{Value: 1}
	header.Next = tail
	node := &LinkedList{Value: 2}
	tail.Next = node
	tail = tail.Next
	node1 := &LinkedList{Value: 3}
	tail.Next = node1
	tail = tail.Next
	node2 := &LinkedList{Value: 5}
	tail.Next = node2
	tail = tail.Next
	node3 := &LinkedList{Value: 5}
	tail.Next = node3
	tail = tail.Next
	node4 := &LinkedList{Value: 6}
	tail.Next = node4
	tail = tail.Next

	// tail.Value := 2
	// tail := tail.Next
	// tail.Value := 3
	// tail := tail.Next
	// tail.Value :=4

	fmt.Println("List := ", header)
	PrintLinkedList(header)
	fmt.Println("Removing Dups : ", RemoveDuplicatesFromLinkedList(header))
	PrintLinkedList(header)

}

func PrintLinkedList(linkedlist *LinkedList) {

	head := linkedlist

	for head.Next != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	head := linkedList
	curr := head
	for curr != nil {
		ptr := curr.Next
		for ptr != nil && ptr.Value == curr.Value {
			ptr = ptr.Next
		}
		curr.Next = ptr
		curr = curr.Next
	}
	return head
}
