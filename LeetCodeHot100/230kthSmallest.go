package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var findk func(root *TreeNode)
	ans := 0
	flag := false
	findk = func(root *TreeNode) {
		if root == nil || flag {
			return
		}
		if root.Left != nil {
			findk(root.Left)
		}
		k--
		if k == 0 && !flag {
			flag = true
			ans = root.Val
			return
		}
		if root.Right != nil {
			findk(root.Right)
		}
	}
	findk(root)
	return ans
}
