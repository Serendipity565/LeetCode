package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := 0
	head := l1
	var pre *ListNode
	for l1 != nil && l2 != nil {
		l1.Val += l2.Val + temp
		if l1.Val >= 10 {
			l1.Val -= 10
			temp = 1
		} else {
			temp = 0
		}
		pre = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		l1.Val += temp
		if l1.Val >= 10 {
			l1.Val -= 10
			temp = 1
		} else {
			temp = 0
		}
		if l1.Next == nil && temp == 1 {
			l1.Next = &ListNode{Val: 1, Next: nil}
			temp = 0
		}
		l1 = l1.Next
	}
	for l2 != nil {
		pre.Next = &ListNode{Val: l2.Val + temp, Next: nil}
		pre = pre.Next
		if pre.Val >= 10 {
			pre.Val -= 10
			temp = 1
		} else {
			temp = 0
		}
		if l2.Next == nil && temp == 1 {
			pre.Next = &ListNode{Val: 1, Next: nil}
			temp = 0
		}
		l2 = l2.Next
	}
	if temp == 1 {
		pre.Next = &ListNode{Val: 1, Next: nil}
	}
	return head
}
