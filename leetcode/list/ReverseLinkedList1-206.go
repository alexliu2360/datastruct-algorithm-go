package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归法
func reverseWithRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseWithRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 迭代法
func reverseWithIter(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var newHead *ListNode = nil
	for cur != nil {
		tmpNode := cur.Next
		cur.Next =  newHead
		newHead = cur
		cur = tmpNode
	}
	return newHead
}
