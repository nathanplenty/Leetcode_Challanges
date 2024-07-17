package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", inorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}), "\nExpected: [1,3,2]")
	fmt.Println("Output 2:", inorderTraversal((*TreeNode)(nil)), "\nExpected: []")
	fmt.Println("Output 3:", inorderTraversal(&TreeNode{Val: 1}), "\nExpected: [1]")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. The number of nodes in the tree is in the range [0, 100].
	b. -100 <= Node.val <= 100
*/

// TreeNode given definition for singly-linked list.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal HEAD COMMENT.
func inorderTraversal(root *TreeNode) []int {
	fmt.Println(root)
	// start with wirst first i from list
	// every row adds a number
	//
	return []int{}
}

/*
Explanation of the function threeSum:

1. TEXT
*/
