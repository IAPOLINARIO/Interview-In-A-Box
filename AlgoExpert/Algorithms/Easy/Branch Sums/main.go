package main

// BinaryTree struct
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func main() {
	tree := NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	BranchSums(tree)
}

// BranchSums sums all the nodes in a binary tree
func BranchSums(root *BinaryTree) []int {

	var result []int

	root.sum(&result, root.Value)

	return result
}

func (tree *BinaryTree) sum(result *[]int, currentSum int) {
	beforeSum := currentSum
	if tree.Left != nil {
		currentSum += tree.Left.Value
		tree.Left.sum(result, currentSum)
	}

	if tree.Right != nil {
		beforeSum += tree.Right.Value
		tree.Right.sum(result, beforeSum)
	}

	if tree.Right == nil && tree.Left == nil {
		*result = append(*result, beforeSum)
	}
}

//NewBinaryTree Constructor
func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}

	tree.Insert(values, 0)

	return tree
}

// Insert Add values to a Binary Tree
func (tree *BinaryTree) Insert(values []int, rootIndex int) *BinaryTree {

	if rootIndex >= len(values) {
		return tree
	}

	val := values[rootIndex]

	queue := []*BinaryTree{tree}

	for len(queue) > 0 {
		var current *BinaryTree
		current, queue = queue[0], queue[1:]

		if current.Left == nil {
			current.Left = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Right)
	}

	tree.Insert(values, rootIndex+1)

	return tree
}
