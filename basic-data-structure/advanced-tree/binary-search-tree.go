package advanced_tree

import (
	"bytes"
	"code_brush/basic-data-structure/queue"
	utils "code_brush/utlis"
	"fmt"
)

// 这个版本的BST存在数据有偏的情况，如果数据有序排列，则退化成链表

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

func generateNode(e interface{}) *Node {
	return &Node{
		e: e,
	}
}

type BinarySearchTree struct {
	root *Node
	size int
}

func NewBST() *BinarySearchTree {
	bis := BinarySearchTree{}
	return &bis
}

// 判断是否为空
func (b *BinarySearchTree) IsEmpty() bool {
	return b.size == 0
}

// 获取BST的大小
func (b *BinarySearchTree) GetSize() int {
	return b.size
}

// 插入节点
func (b *BinarySearchTree) InsertNode(e interface{}) {
	b.root = b.insert(b.root, e)
}

// 删除节点
func (b *BinarySearchTree) DeleteNode(e interface{}) {
	b.root = b.deledeNode(b.root, e)
}

// 查看是否包含元素
func (b *BinarySearchTree) IsContain(e interface{}) bool {
	return b.contain(b.root, e)
}

// 查找最小值
func (b *BinarySearchTree) Minimum() interface{} {
	if b.size == 0 {
		fmt.Println("bst size is 0")
		return nil
	}
	return b.minimum(b.root).e
}

// 查找最小值
func (b *BinarySearchTree) Maximum() interface{} {
	if b.size == 0 {
		fmt.Println("bst size is 0")
		return nil
	}
	return b.maximum(b.root).e
}

// 中序遍历
func (b *BinarySearchTree) InOrder() {
	b.inOrder(b.root)
}

// 前序遍历
func (b *BinarySearchTree) PreOrder() {
	b.preOrder(b.root)
}

// 后序遍历
func (b *BinarySearchTree) PostOrder() {
	b.postOrder(b.root)
}

// 层序遍历
func (b *BinarySearchTree) LevelOrder() {
	q := queue.NewLoopQueue(20)
	q.Enqueue(b.root)
	for !q.IsEmpty() {
		cur := q.Dequeue().(*Node)
		fmt.Println(cur.e)

		if cur.left != nil {
			q.Enqueue(cur.left)
		}
		if cur.right != nil {
			q.Enqueue(cur.right)
		}
	}
}

// 删除最小值
func (b *BinarySearchTree) RemoveMin() interface{} {
	// 获得最小值
	ret := b.Minimum()
	b.root = b.removeMin(b.root)
	return ret
}

func (b *BinarySearchTree) RemoveMax() interface{} {
	ret := b.Maximum()
	b.root = b.removeMax(b.root)
	return ret
}

func (b *BinarySearchTree) removeMin(root *Node) *Node {
	if root.left == nil {
		rightNode := root.right // 记录右节点，因为右节点需要上移
		root.right = nil  // 将右节点设为空，不管原先有值还是为nil
		b.size--
		return rightNode // 返回右节点
	}
	root.left = b.removeMin(root.left) // 将上次上移的右节点赋值（替换）root的左节点
	return root
}

func (b *BinarySearchTree) String() string {
	var buffer bytes.Buffer
	b.generateBSTString(b.root, 0, &buffer)
	return buffer.String()
}
func (b *BinarySearchTree) insert(root *Node, e interface{}) *Node {
	if root == nil {
		b.size++
		return generateNode(e)
	}

	if utils.Compare(e, root.e) < 0 {
		root.left = b.insert(root.left, e)
	} else if utils.Compare(e, root.e) > 0 {
		root.right = b.insert(root.right, e)
	}
	return root
}

func (b *BinarySearchTree) contain(root *Node, e interface{}) bool {
	if root == nil {
		return false
	}
	if utils.Compare(e, root.e) == 0 {
		return true
	} else if utils.Compare(root.left, e) < 0 {
		return b.contain(root.left, e)
	} else {
		return b.contain(root.right, e)
	}
}

func (b *BinarySearchTree) preOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.e)
	b.preOrder(root.left)
	b.preOrder(root.right)
}

func (b *BinarySearchTree) generateBSTString(root *Node, depth int, buffer *bytes.Buffer) {
	if root == nil {
		buffer.WriteString(b.generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(b.generateDepthString(depth) + fmt.Sprint(root.e) + "\n")
	b.generateBSTString(root.left, depth+1, buffer)
	b.generateBSTString(root.right, depth+1, buffer)
}

func (b *BinarySearchTree) generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}

func (b *BinarySearchTree) inOrder(root *Node) {
	if root == nil {
		return
	}

	b.inOrder(root.left)
	fmt.Println(root.e)
	b.inOrder(root.right)
}

func (b *BinarySearchTree) postOrder(root *Node) {
	if root == nil {
		return
	}

	b.postOrder(root.left)
	b.postOrder(root.right)
	fmt.Println(root.e)
}

func (b *BinarySearchTree) minimum(root *Node) *Node {
	if root.left == nil {
		return root
	}
	return b.minimum(root.left)
}

func (b *BinarySearchTree) maximum(root *Node) *Node {
	if root.right == nil {
		return root
	}

	return b.maximum(root.right)
}

/////////////////
//      5      //
//    /   \    //
//   3    8    //
//  / \  /     //
// 2  4  6     //
/////////////////
func (b *BinarySearchTree) removeMax(root *Node) *Node {
	if root.right == nil{
		leftNode := root.left
		root.left = nil
		b.size--
		return leftNode
	}

	root.right = b.removeMax(root.right)
	return root
}

func (b *BinarySearchTree) deledeNode(root *Node, e interface{}) *Node {
	if root == nil{
		fmt.Printf("the '%+v' is not in BST\n", e)
		return nil
	}

	if utils.Compare(e, root.e) < 0{
		root.left = b.deledeNode(root.left, e)
		return root
	} else if utils.Compare(e, root.e) > 0{
		root.right = b.deledeNode(root.right, e)
		return root
	} else {
		// left == nil
		// 上移该节点
		if root.left == nil{
			rightNode := root.right
			root.right = nil
			b.size--
			return rightNode
		}
		// right == nil
		// 上移该节点
		if root.right == nil{
			leftNode := root.left
			root.right = nil
			b.size--
			return leftNode
		}
		// left!=nil && right!=nil
		// 此时从待删除节点的右子树上寻找最小节点
		// 用寻找到的最小节点替代被删除的节点
		successor := b.minimum(root.right) // 根据BST的特点，后继者为最小值，寻找后继节点
		successor.right = b.removeMin(root.right) //  （右节点）用寻找到的最小节点替代被删除的节点，并删除原有的最小值
		successor.left = root.left //（左节点）

		root.right = nil
		root.left = nil

		return successor
	}
}


// floor

// cell

