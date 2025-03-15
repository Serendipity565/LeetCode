package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	mergeTwo := func(list1 *ListNode, list2 *ListNode) *ListNode {
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
	ans := lists[0]
	for i := 0; i < len(lists)-1; i++ {
		ans = mergeTwo(ans, lists[i+1])
	}
	return ans
}
