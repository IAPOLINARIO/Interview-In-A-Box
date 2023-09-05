package main

import "fmt"

// BST struct to store tree values
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

//NewBST is a Constructor for a BST tree
func NewBST(value int) *BST {
	return &BST{Value: value}
}

func main() {

	tree := NewBST(10)
	tree.Left = NewBST(5)
	tree.Left.Left = NewBST(2)
	tree.Left.Left.Left = NewBST(1)
	tree.Left.Right = NewBST(5)
	tree.Right = NewBST(15)
	tree.Right.Left = NewBST(13)
	tree.Right.Left.Right = NewBST(14)
	tree.Right.Right = NewBST(22)

	valueToFind := 12

	fmt.Print(tree.FindClosestValue(valueToFind))

}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

// FindClosestValue finds the closest value in a BST
func (tree *BST) FindClosestValue(target int) int {
	closestValue := tree.Value

	for tree != nil {

		if abs(tree.Value-target) < abs(closestValue-target) {
			closestValue = tree.Value
		}

		if target < tree.Value {
			tree = tree.Left
		} else {
			tree = tree.Right
		}

	}

	return closestValue

}
