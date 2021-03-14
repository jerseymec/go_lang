package main

import "fmt"

//parent index * 2 +1 = left child index
//parent index *2 +2 = right child index
//max heap is built on a slice that holds the array
type MaxHeap struct {
	array []int
}

//Insert adds an element to the heap
func (h *MaxHeap) Insert(Key int) {
	h.array = append(h.array, Key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// newly inserted node at the right lowest level is put in the root
//MaxHeapifyUp will heapify from bottom up
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}

}
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	// for loop to traverse down the path till node have atleast one child
	for l <= lastIndex {
		//when l is the only child
		// when l is the greater than r child
		if l == lastIndex || h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			//else compare to r child
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)

		} else {
			return
		}

	}
}

//get parent index from the left child index
func parent(ind int) int {
	return (ind - 1) / 2
}

//get the left child indez
func left(i int) int {
	return i*2 + 1
}

//get the right child indez
func right(i int) int {
	return i*2 + 2
}

//swap keys in the heap slice
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

//Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	if l < 0 {
		err := fmt.Errorf("Heap is Empty")
		fmt.Println(err.Error())
		return -1
	}
	//put the last node in heap to the root of tree
	h.array[0] = h.array[l]
	//reduce the size of the tree by 1
	h.array = h.array[:l]
	//now heapify tree from root down to find correct position of current root
	h.maxHeapifyDown(0)
	return extracted
}

// This func will heapify from top down

func main() {
	//Maxpriority := &MaxHeap{[] int{10,20,30}}
	Maxpriority := &MaxHeap{}
	build1 := []int{10, 20, 30}
	for _, v := range build1 {
		Maxpriority.Insert(v)
		fmt.Println(Maxpriority)
	}

	fmt.Println(Maxpriority)
	buildheap := []int{21, 35, 60, 80}

	for _, v := range buildheap {
		Maxpriority.Insert(v)
		fmt.Println(Maxpriority)
	}
	fmt.Println(Maxpriority.Extract())
	fmt.Println(Maxpriority)

}
