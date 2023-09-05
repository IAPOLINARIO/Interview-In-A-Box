package main

import "fmt"

// Node is a struct
type Node struct {
	Name     string
	Children []*Node
}

//NewNode is a constructor
func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

//AddChildren adds a new node
func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		child := Node{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}

func main() {
	var graph = NewNode("A").AddChildren("B", "C", "D")
	graph.Children[0].AddChildren("E").AddChildren("F")
	graph.Children[2].AddChildren("G").AddChildren("H")
	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
	graph.Children[2].Children[0].AddChildren("K")

	output := graph.DepthFirstSearch([]string{})

	fmt.Print(output)
}

//DepthFirstSearch is a extension method for the Node struct
func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, node := range n.Children {
		array = append(array, node.Name)
		for _, granchildNode := range node.Children {
			array = granchildNode.DepthFirstSearch(array)
		}

	}
	return array
}
