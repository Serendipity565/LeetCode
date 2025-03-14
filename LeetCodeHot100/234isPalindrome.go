package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	lenth := func(head *ListNode) int {
		length := 0
		for head != nil {
			length++
			head = head.Next
		}
		return length
	}
	rever := func(head *ListNode) *ListNode {
		var pre *ListNode
		curr := head
		for curr != nil {
			next := curr.Next
			curr.Next = pre
			pre = curr
			curr = next
		}
		return pre
	}

	lenHead := lenth(head)
	midLen := (lenHead + 1) / 2
	midHead := head
	for i := 0; i < midLen; i++ {
		midHead = midHead.Next
	}
	midHead = rever(midHead)
	for midHead != nil {
		if head.Val != midHead.Val {
			return false
		}
		head = head.Next
		midHead = midHead.Next
	}
	return true
}
