package LeetCodeHot100

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var check func(root *TreeNode, mymin int, mymax int) bool
	check = func(root *TreeNode, mymin int, mymax int) bool {
		if root == nil {
			return true
		} else if root.Val < mymax && root.Val > mymin {
			return check(root.Left, mymin, root.Val) && check(root.Right, root.Val, mymax)
		} else {
			return false
		}
	}
	return check(root, math.MinInt64, math.MaxInt64)
}
