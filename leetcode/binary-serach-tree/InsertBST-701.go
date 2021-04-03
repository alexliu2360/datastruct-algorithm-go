package binary_serach_tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return &TreeNode{
			Val:val,
			Right:nil,
			Left:nil,
		}
	}

	if root.Val < val{
		root.Right = insertIntoBST(root.Right,val)
	}

	if root.Val > val{
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}