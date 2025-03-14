package LeetCodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = 1
		head = head.Next
	}
	return nil
}
