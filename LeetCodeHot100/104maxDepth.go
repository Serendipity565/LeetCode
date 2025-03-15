package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var deep func(*TreeNode) int
	deep = func(root *TreeNode) int {
		if root.Right == nil && root.Left == nil {
			return 1
		}
		if root.Left != nil && root.Right != nil {
			return max(deep(root.Left), deep(root.Right)) + 1
		} else if root.Right != nil {
			return deep(root.Right) + 1
		} else {
			return deep(root.Left) + 1
		}
	}
	if root == nil {
		return 0
	} else {
		return deep(root)
	}
}
