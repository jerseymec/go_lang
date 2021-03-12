package main

import "fmt"

type Stack struct {
	items []int
}

//Push

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//pop will remove a value at the end
// and returns the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:len(s.items)-1]
	return toRemove

}

func main() {
	myStack := Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	myStack.Push(600)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)

	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)

	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
