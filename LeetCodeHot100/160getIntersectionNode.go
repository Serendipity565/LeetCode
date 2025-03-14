package LeetCodeHot100

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	getLen := func(head *ListNode) int {
		length := 0
		for head != nil {
			length++
			head = head.Next
		}
		return length
	}

	lenA := getLen(headA)
	lenB := getLen(headB)

	// 对齐长度
	for lenA > lenB {
		headA = headA.Next
		lenA--
	}
	for lenB > lenA {
		headB = headB.Next
		lenB--
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	if headA == nil {
		return nil
	} else {
		return headA
	}
}
