package LeetCodeHot100

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		head = next
	}
	fmt.Println(head.Val)

	swap := func(pre *ListNode, cur *ListNode, next *ListNode) {
		cur.Next = next.Next
		pre.Next = next
		next.Next = cur
	}
	pre := head.Next
	// fmt.Println(pre.Val)
	for pre.Next != nil && pre.Next.Next != nil {
		cur := pre.Next
		// fmt.Println(cur.Val)
		next := cur.Next
		if next == nil {
			return head
		}
		swap(pre, cur, next)
		pre = pre.Next.Next
	}
	return head
}
