package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := -1000
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lVal := dfs(node.Left)
		rVal := dfs(node.Right)
		ans = max(ans, lVal+rVal+node.Val)
		return max(max(lVal, rVal)+node.Val, 0)
	}
	dfs(root)
	return ans
}
