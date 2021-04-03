package set

import (
	"code_brush/basic-data-structure/advanced-tree"
)

//基于bst的set数据结构

type BstSet struct {
	Bst *advanced_tree.BinarySearchTree
}

func NewBstSet() *BstSet{
	return &BstSet{
		Bst: advanced_tree.NewBST(),
	}
}

func (b *BstSet) Add(e interface{}) {
	b.Bst.InsertNode(e)
}

func (b *BstSet) Remove(e interface{}) {
	b.Bst.DeleteNode(e)
}

func (b *BstSet) Contains(e interface{}) bool {
	return b.Bst.IsContain(e)
}


func (b *BstSet) GetSize() int {
	return b.Bst.GetSize()
}

func (b *BstSet) IsEmpty() bool {
	return b.Bst.IsEmpty()
}
