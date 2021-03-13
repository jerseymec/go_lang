package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	Key       int
	Adjacency []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{Key: k})
	}
}

//Adds an edge to the graph
func (g *Graph) AddEdge(fromKey, toKey int) {
	fromVertex := g.getVertex(fromKey)
	toVertex := g.getVertex(toKey)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v -> %v) ", fromKey, toKey)
		fmt.Println(err.Error())
	} else if contains(fromVertex.Adjacency, toKey) {
		err := fmt.Errorf("Existing Edge (%v -> %v) ", fromKey, toKey)
		fmt.Println(err.Error())
	} else {
		fromVertex.Adjacency = append(fromVertex.Adjacency, toVertex)
	}

}

//returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.Key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, k int) bool {
	for _, i := range s {
		if k == i.Key {
			return true
		}
	}
	return false

}

func (g Graph) PrintGraph() {
	for _, i := range g.vertices {
		fmt.Printf("\n Vertex: %v", i.Key)
		for _, vert := range i.Adjacency {
			fmt.Printf(" Edges: %v", vert.Key)
		}
	}
	fmt.Println()
}

func main() {
	DFS := &Graph{}

	for i := 0; i <= 5; i++ {
		DFS.AddVertex(i)
		fmt.Println(DFS)

	}
	DFS.AddVertex(0)
	DFS.AddVertex(0)

	fmt.Println(" Printing all Vertices ")
	DFS.PrintGraph()
	DFS.AddEdge(6, 2)
	DFS.AddEdge(5, 2)
	DFS.AddEdge(5, 2)
	DFS.PrintGraph()

}
