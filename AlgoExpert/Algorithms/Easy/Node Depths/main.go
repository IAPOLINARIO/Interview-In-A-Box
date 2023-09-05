package main

import "fmt"

//BinaryTree struct
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func main() {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Left.Left = &BinaryTree{Value: 8}
	root.Left.Left.Right = &BinaryTree{Value: 9}
	root.Left.Right = &BinaryTree{Value: 5}
	root.Right = &BinaryTree{Value: 3}
	root.Right.Left = &BinaryTree{Value: 6}
	root.Right.Right = &BinaryTree{Value: 7}

	NodeDepths(root)
}

//NodeDepths calculates the depth of a given binary tree
func NodeDepths(root *BinaryTree) int {

	if root == nil {
		return 0
	}

	depth := 0

	if root.Left != nil {
		depth += calculateDepth(root.Left, 1)
	}

	if root.Right != nil {
		depth += calculateDepth(root.Right, 1)
	}

	fmt.Print(depth)
	return depth
}

func calculateDepth(node *BinaryTree, currentDepth int) int {
	fmt.Printf("Value: %d - Depth: %d \n", node.Value, currentDepth)

	if node.Left == nil && node.Right == nil {
		return currentDepth
	}

	finalDepth := currentDepth //1

	if node.Left != nil {
		finalDepth += calculateDepth(node.Left, currentDepth+1) //2
	}

	if node.Right != nil {
		finalDepth += calculateDepth(node.Right, currentDepth+1)
	}

	//finalDepth = finalDepth + currentDepth

	return finalDepth
}
