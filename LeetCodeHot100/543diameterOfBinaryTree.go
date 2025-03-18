package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var deep func(root *TreeNode) int
	ans := 0
	deep = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Right == nil && root.Left == nil {
			return 1
		}
		left := deep(root.Left)
		right := deep(root.Right)
		ans = max(ans, left+right+1)
		return max(left, right) + 1
	}

	deep(root)
	return ans - 1
}
