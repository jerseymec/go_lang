package main

import "fmt"

func QuickSort(array []int) []int {
	// Write your code here.
	return QS(array,0,len(array)-1)
}

func QS(array []int, startIdx,endIdx int) []int {
	if startIdx >= endIdx{
		return array
	}
	pivotIdx := startIdx
	leftIdx := startIdx +1
	rightIdx := endIdx
	for leftIdx <= rightIdx{
		switch{
		case array[leftIdx] > array[pivotIdx] && array[rightIdx] < array[pivotIdx]:
			swap(leftIdx,rightIdx,array)

		case array[leftIdx] <= array[pivotIdx]:
			leftIdx++
		case array[rightIdx] >=array[pivotIdx]:
			rightIdx--
		}
	}
	swap(pivotIdx,rightIdx,array)
	switch rightIdx-1-startIdx < endIdx- (rightIdx+1){
	case true:
		QS(array,startIdx,rightIdx-1)
		QS(array,rightIdx+1,endIdx)
	case false:
		QS(array,rightIdx+1,endIdx)
		QS(array,startIdx, rightIdx-1)


	}

	return array
}

func swap(i,j int, array []int){
	array[i],array[j] = array[j],array[i]
}


func main() {
	seq := []int{2, 4, 52, 5, 8, 9}
	fmt.Println("Sort of [2,4,52,5,8,9] = ", QuickSort(seq))
}
