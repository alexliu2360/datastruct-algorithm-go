package binary_serach_tree

import "fmt"
import "gopkg.in/eapache/queue.v1"
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else {
		if root.Left == nil {
			rightNode := root.Right
			root.Right = nil
			return rightNode
		}

		if root.Right == nil {
			leftNode := root.Left
			root.Left = nil
			return leftNode
		}

		successor := getMin(root.Right)
		successor.Right = removeMin(root.Right)
		successor.Left = root.Left


		root.Left = nil
		root.Right = nil

		return successor
	}
}

func getMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	return getMin(root.Left)
}

func removeMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		rightNode := root.Right
		root.Right = nil
		return rightNode
	}

	root.Left = removeMin(root.Left)
	return root
}


func insertNode(root *TreeNode, e int)*TreeNode{
	if root == nil{
		return &TreeNode{
			Val:   e,
			Left:  nil,
			Right: nil,
		}
	}
	if root.Val > e{
		root.Left = insertNode(root.Left, e)
	}
	if root.Val < e{
		root.Right = insertNode(root.Right, e)
	}
	return root
}

func levelPrint(root *TreeNode){
	q := queue.New()
	q.Add(root)
	for 0 != q.Length() {
		cur := q.Remove().(*TreeNode)

		fmt.Println(cur.Val)
		if cur.Left != nil{
			q.Add(cur.Left)
		}
		if cur.Right != nil{
			q.Add(cur.Right)
		}
	}
}

func printBST(root *TreeNode){
	if root == nil{
		return
	}
	printBST(root.Left)
	fmt.Println(root.Val)
	printBST(root.Right)
}