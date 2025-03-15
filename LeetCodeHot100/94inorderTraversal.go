package LeetCodeHot100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var traversal func(*TreeNode)
	traversal = func(root *TreeNode) {
		if root.Left != nil {
			traversal(root.Left)
		}
		ans = append(ans, root.Val)
		if root.Right != nil {
			traversal(root.Right)
		}
	}
	if root != nil {
		traversal(root)
	}
	return ans
}
