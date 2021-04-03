package binary_serach_tree

/*
假设一个二叉搜索树具有如下特征：
	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。
*/

func validateBST(root *TreeNode, min *TreeNode, max *TreeNode)bool{
	if root == nil{return true}
	if min != nil && (root.Val ==  min.Val || root.Val <  min.Val) {return false}
	if max != nil && (root.Val ==  max.Val || root.Val >  max.Val) {return false}

	return validateBST(root.Left, min, root) && validateBST(root.Right, root, max)
}

func isValidBST(root *TreeNode) bool {
	return validateBST(root, nil, nil)
}
