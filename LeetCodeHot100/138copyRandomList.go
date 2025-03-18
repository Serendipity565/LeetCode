package LeetCodeHot100

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	var newHead *Node
	if head == nil {
		return newHead
	}
	// 虚拟头节点
	newHead = &Node{
		Val:    0,
		Next:   nil,
		Random: nil,
	}
	ans := newHead
	// 深拷贝节点
	for p := head; p != nil; p = p.Next {
		newHead.Next = &Node{
			Val:    p.Val,
			Next:   nil,
			Random: p.Random,
		}
		newHead = newHead.Next
	}
	// 处理random指针
	newHead = ans
	for p := head; p != nil; p = p.Next {
		if p.Random == nil {
			newHead.Next.Random = nil
		} else {
			newRadnom := ans.Next
			for q := head; q != p.Random; q = q.Next {
				q = q.Next
				newRadnom = newRadnom.Next
			}
			newHead.Next.Random = newRadnom
		}

		newHead = newHead.Next
	}
	return ans.Next
}
