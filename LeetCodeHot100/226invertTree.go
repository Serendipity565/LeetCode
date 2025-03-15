package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var reverse func(*TreeNode)
	reverse = func(root *TreeNode) {
		root.Left, root.Right = root.Right, root.Left
		if root.Left != nil {
			reverse(root.Left)
		}
		if root.Right != nil {
			reverse(root.Right)
		}
	}
	reverse(root)
	return root
}
