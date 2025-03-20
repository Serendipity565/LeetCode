package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var res [][]int
	var virtualRoot = &TreeNode{Val: -1, Left: root, Right: nil}

	var level func(root *TreeNode, deep int)
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
	var ans []int
	for _, v := range res {
		ans = append(ans, v[len(v)-1])
	}
	return ans
}
