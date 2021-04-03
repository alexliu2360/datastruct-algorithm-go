package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ===========================================================迭代法
func reverseNWithIter(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var newHead *ListNode = nil
	for i := 0; i < n && cur != nil; i++ {
		tmpNode := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = tmpNode
	}
	head.Next = cur
	return newHead
}

func reverseBetweenWithIter(head *ListNode, left int, right int) *ListNode {
	if left > right {
		return nil
	}
	if head == nil || head.Next == nil {
		return head
	}
	//或者设置哨兵头节点，dummyHead，返回的时候，返回dummyHead.Next即可
	var preHead *ListNode = nil
	cur := head
	// 先指向left
	for i := 0; i < left - 1 && cur != nil; i++ {
		preHead = cur
		cur = cur.Next
	}

	var newHead *ListNode = nil
	var newCur = cur
	for i := left-1; i < right && cur != nil; i++ {
		tmpNode := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = tmpNode
	}
	if newCur != nil && cur != nil{
		newCur.Next = cur
	}
	if preHead != nil {
		preHead.Next = newHead
	} else {
		head = newHead
	}

	return head
}

// ===========================================================递归法
var successor *ListNode = nil

func reverseNWithRecursion(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseNWithRecursion(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

func reverseBetweenWithRecursion(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseNWithRecursion(head, right)
	}

	head.Next = reverseBetweenWithRecursion(head.Next, left-1, right-1)
	return head
}
