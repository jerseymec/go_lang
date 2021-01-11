package main

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {

	array = append(array, n.Name)
	for _, i := range n.Children {
		array = i.DepthFirstSearch(array)
	}
	return array
}

func main(){
	
}