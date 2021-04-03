package list

import (
	"fmt"
	"testing"
)

func genList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := ListNode{
		Val:  arr[0],
		Next: nil,
	}
	next := &head
	for i := 1; i < len(arr); i++ {
		ln := ListNode{
			Val:  arr[i],
			Next: nil,
		}
		next.Next = &ln
		next = &ln
	}
	return &head
}

func printList(l *ListNode, n int){
	if l == nil{
		return
	}
	next := l
	fmt.Printf("[%d] ", n)
	for next != nil {
		fmt.Printf("%d -> ", next.Val)
		if next.Next == nil{
			fmt.Println("nil")
		}
		next = next.Next
	}
}

func TestReverseWithIter(t *testing.T) {
	//arr := []int{2, 4, 1, 8, 6,9,10}
	arr1 := []int{5}

	l1 := genList(arr1)
	printList(l1,1)
	l1 = reverseBetweenWithIter(l1, 1,1)
	printList(l1,2)
	//l := genList(arr)
	//printList(l,1)
	////l = reverseNWithIter(l, 5)
	//l = reverseBetweenWithIter(l, 1,3)
	//printList(l,2)
	//l = reverseBetweenWithIter(l, 3,6)
	//printList(l,2)
	//l = reverseBetweenWithIter(l, 1,7)
	//printList(l,2)
	//l = reverseBetweenWithIter(l, 3,7)
	//printList(l,2)
	//l = reverseWithRecursion(l)
	//printList(l, 2)
	//l = reverseWithIter(l)
	//printList(l,3)
	//l = reverseNWithRecursion(l, 3)
	//printList(l,4)
	//l = reverseBetweenWithRecursion(l, 2,4)
	//printList(l,5)
}
