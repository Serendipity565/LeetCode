package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	temp := &ListNode{Val: 0, Next: head}
	reverse := func(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
		pre := head
		cur := head.Next
		for cur != tail {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		cur.Next = pre
		return tail, head
	}

	pre := temp
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return temp.Next
			}
		}
		nex := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}

	return temp.Next
}
