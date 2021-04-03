package advanced_tree

// 线段树，是平衡二叉树，不是完全二叉树

type SegmentTree struct {
	tree []interface{}
	data []interface{}
}

func NewSegementTree(arr []interface{}) *SegmentTree {
	st := &SegmentTree{
		tree: make([]interface{}, len(arr)*4),
		data: make([]interface{}, len(arr)),
	}
	for i := 0; i < len(arr); i++ {
		st.data[i] = arr[i]
	}

	return st
}

func (st *SegmentTree) GetSize() int {
	return len(st.data)
}
