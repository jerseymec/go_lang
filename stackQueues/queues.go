package main

//Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

//Enqueue adds a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue
func (q *Queue) Dequeue() int {

	popped := q.items[0]
	q.items = q.items[1:]
	return popped
}
