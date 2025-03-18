package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var level func(root *TreeNode, deep int)
	var res [][]int
	virtualRoot := &TreeNode{Val: -1, Left: root, Right: nil}

	level = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}
		if root.Left != nil {
			if deep+1 >= len(res) {
				res = append(res, []int{})
			}
			res[deep+1] = append(res[deep+1], root.Left.Val)
			level(root.Left, deep+1)
		}
		if root.Right != nil {
			if deep+1 >= len(res) {
				res = append(res, []int{})
			}
			res[deep+1] = append(res[deep+1], root.Right.Val)
			level(root.Right, deep+1)
		}
		return
	}

	level(virtualRoot, -1)
	return res
}
