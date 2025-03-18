package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMiddle(head)
	right := mid.Next
	mid.Next = nil

	leftSorted := sortList(head)
	rightSorted := sortList(right)

	return mergeTwo(leftSorted, rightSorted)
}

func getMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwo(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var ans *ListNode
	var temp *ListNode

	ans = &ListNode{Val: -1, Next: nil}
	temp = ans
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ans.Next = list1
			list1 = list1.Next
		} else {
			ans.Next = list2
			list2 = list2.Next
		}
		ans = ans.Next
	}
	for list1 != nil {
		ans.Next = list1
		list1 = list1.Next
		ans = ans.Next
	}
	for list2 != nil {
		ans.Next = list2
		list2 = list2.Next
		ans = ans.Next
	}
	return temp.Next
}
