package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	var check func(left, right *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		} else {
			return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
		}
	}

	if root == nil {
		return true
	} else {
		return check(root.Left, root.Right)
	}
}
