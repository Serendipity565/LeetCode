package LeetCodeHot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var digui func(left, right int) *TreeNode
	digui = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{Val: nums[mid], Left: nil, Right: nil}
		root.Left = digui(left, mid-1)
		root.Right = digui(mid+1, right)
		return root
	}
	return digui(0, len(nums)-1)
}
