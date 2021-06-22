package main

func main() {

}
func MinHeightBST(array []int) *BST {

	return constructbst(array, 0, len(array)-1)

}

func constructbst(array []int, startindex, endindex int) *BST {
	if startindex > endindex {
		return nil
	}
	midIndex := (startindex + endindex) / 2
	newRoot := &BST{array[midIndex]}
	newRoot.Left = construcbst(array, startindex, midIndex-1)
	newRoot.Right = construcbst(array, midIndex+1, endindex)
	return newRoot
}

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
