package binary_serach_tree

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	x := []int{8,3,1,5,9}
	var r *TreeNode
	for _, i :=range x{
		r = insertNode(r,i)
	}
	printBST(r)
	deleteNode(r, 1)
	fmt.Println("---------")

	printBST(r)
	fmt.Println("---------")
	levelPrint(r)
}

func TestInsertNode(t *testing.T) {
	x := []int{8,3,1,5,9}
	var r *TreeNode
	for _, i :=range x{
		r = insertIntoBST(r,i)
	}
	printBST(r)
	fmt.Println("---------")
	levelPrint(r)
}

func TestIsSame(t *testing.T) {
	t1 := &TreeNode{
		Val:1,
	}
	t2 := &TreeNode{
		Val:2,
	}

	t3 := &TreeNode{
		Val:2,
		Left:t1,
		Right:t2,
	}
	t4 := &TreeNode{
		Val:2,
		Left:t1,
		Right:t2,
	}
	fmt.Println(isSameTree(t1, t2))
	fmt.Println(isSameTree(t3, t4))
}


func TestIsValidBST(t *testing.T) {
	t1 := &TreeNode{
		Val:4,
	}
	t2 := &TreeNode{
		Val:10,
	}

	t3 := &TreeNode{
		Val:6,
		Left:t1,
		Right:t2,
	}
	t4 := &TreeNode{
		Val:3,
		Left:t1,
		Right:t2,
	}
	fmt.Println(isValidBST(t3))
	fmt.Println(isValidBST(t4))

}