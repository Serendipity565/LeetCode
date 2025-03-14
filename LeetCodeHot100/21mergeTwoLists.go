package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var head *ListNode
	var temp *ListNode

	head = &ListNode{Val: -1, Next: nil}
	temp = head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	for list1 != nil {
		head.Next = list1
		list1 = list1.Next
		head = head.Next
	}
	for list2 != nil {
		head.Next = list2
		list2 = list2.Next
		head = head.Next
	}
	return temp.Next
}
