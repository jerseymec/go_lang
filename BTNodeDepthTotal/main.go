package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (root *BST) Nodedepth(depth int) int {

	return root.printdepth(0)
}
func (root *BST) printdepth(totaldepth int) int {

	if root == nil {
		return 0
	}
	fmt.Println("Node with Value ={root.value} has depth {totaldepth} ")

	return totaldepth + root.Left.printdepth(totaldepth+1) + root.Right.printdepth(totaldepth+1)
}

func main() {

}
