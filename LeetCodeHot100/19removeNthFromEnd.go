package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lenth := func(head *ListNode) int {
		count := 0
		for head != nil {
			count++
			head = head.Next
		}
		return count
	}

	lenH := lenth(head) - n
	if lenH == 0 {
		return head.Next
	}

	temp := head
	var pre *ListNode
	for lenH > 0 {
		pre = temp
		temp = temp.Next
		lenH--
	}
	if temp.Next == nil {
		pre.Next = nil
	} else {
		pre.Next = temp.Next
	}
	return head
}
