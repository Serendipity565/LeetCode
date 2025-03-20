package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	var change func(root *TreeNode) *TreeNode
	change = func(root *TreeNode) *TreeNode {
		if root == nil || (root.Right == nil && root.Left == nil) {
			return root
		}
		if root.Left == nil && root.Right != nil {
			return change(root.Right)
		}
		if root.Left != nil && root.Right == nil {
			root.Right = root.Left
			root.Left = nil
			return change(root.Right)
		}
		temp := root.Right
		root.Right = root.Left
		root.Left = nil
		change(root.Right).Right = temp
		return change(temp)
	}
	change(root)
}
