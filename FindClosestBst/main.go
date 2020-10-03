package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func absdiff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	closest := tree.Value
	return tree.traverse(target, closest)
}
func (root *BST) traverse(searchval int, closest int) int {

	x1 := absdiff(root.Value, searchval)
	x2 := absdiff(closest, searchval)
	if x1 < x2 {
		closest = root.Value
	}
	if root.Value > searchval && root.Left != nil {
		return root.Left.traverse(searchval, closest)

	} else if root.Value < searchval && root.Right != nil {
		return root.Right.traverse(searchval, closest)
	}

	return closest
}

func main() {

}
