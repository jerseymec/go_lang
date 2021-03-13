package main

import "fmt"

type Node struct {
	Key  int
	Next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

// func (n *LinkedList)Initialize (first *Node) *LinkedList{
// 	n := &LinkedList{first}
// 	n.length := 1
// 	return l
// }

func (l *LinkedList) Insert(n *Node) {
	second := l.head
	l.head = n
	l.head.Next = second
	l.length++
}

func (l *LinkedList) deleteWithValue(val int) {
	if l.length == 0 {
		return
	}
	if l.head.Key == val {
		l.head = l.head.Next
		l.length--
	}
	previousToDelete := l.head

	for previousToDelete.Next.Next != nil && previousToDelete.Next.Key != val {
		previousToDelete = previousToDelete.Next

	}
	if previousToDelete.Next == nil {
		return
	}
	if previousToDelete.Next.Key == val {
		previousToDelete.Next = previousToDelete.Next.Next
		l.length--
	}
	return
}

func (l LinkedList) printListData() {
	leng := l.length
	curr := l.head
	for leng != 0 {

		fmt.Println(curr.Key)
		leng--
		curr = curr.Next
	}

}

func main() {
	n := &Node{Key: 100}
	n1 := &Node{Key: 200}
	n2 := &Node{Key: 300}
	n3 := &Node{Key: 400}
	fmt.Println(n)
	ll := &LinkedList{}
	fmt.Println(ll)
	ll.Insert(n)
	ll.Insert(n1)
	ll.Insert(n2)
	ll.Insert(n3)
	fmt.Println(ll)
	ll.printListData()
	ll.deleteWithValue(200)
	fmt.Println("-----------")
	ll.printListData()

}
