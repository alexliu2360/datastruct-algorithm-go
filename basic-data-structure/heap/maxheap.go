package heap

import (
	"code_brush/basic-data-structure/array"
	utils "code_brush/utlis"
	"fmt"
)

// 最大（小）堆是指在树中，存在一个结点而且该结点有儿子结点，该结点的data域值都不小于（大于）其儿子结点的data域值，并且它是一个完全二叉树（不是满二叉树）
// 本例是基于array的最大堆
type MaxHeap struct {
	data *array.Array
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: array.New(20), // default 20
	}
}

func (h *MaxHeap) Size() int {
	return h.data.GetSize()
}

func Heapify(arr []interface{}) *MaxHeap {
	maxHeap := MaxHeap{data: array.GetArrayFromArr(arr)}
	for i := maxHeap.parentIndex(len(arr) - 1); i >= 0; i-- {
		maxHeap.ShiftDown(i)
	}
	return &maxHeap
}

func (h *MaxHeap) IsEmpty() bool {
	return h.data.IsEmpty()
}
func (h *MaxHeap) FindMax() interface{} {
	if h.data.GetSize() == 0 {
		fmt.Println("error: the heap is empty!")
		return nil
	}
	return h.data.Get(0)
}

func (h *MaxHeap) Add(e interface{}) {
	h.data.AddLast(e)
	h.ShiftUp(h.data.GetSize() - 1)
}

func (h *MaxHeap) Replace(e interface{}) interface{} {
	ret := h.FindMax()
	h.data.Set(0, e)
	h.ShiftDown(0)
	return ret
}

func (h *MaxHeap) ExtractMax() interface{} {
	ret := h.FindMax()
	if ret != nil {
		h.data.Swap(0, h.data.GetSize()-1)
		h.data.RemoveLast()
		h.ShiftDown(0)
	}
	return ret
}
func (h *MaxHeap) parentIndex(index int) int {
	if index == 0 {
		fmt.Println("error: index is 0")
		panic("index-0 doesn't have parentIndex")
	}

	return (index - 1) / 2
}

func (h *MaxHeap) leftChildIndex(index int) int {
	return index*2 + 1
}

func (h *MaxHeap) rightChildIndex(index int) int {
	return index*2 + 2
}

func (h *MaxHeap) ShiftUp(k int) {
	for k > 0 && utils.Compare(h.data.Get(k), h.data.Get(h.parentIndex(k))) > 0 {
		h.data.Swap(k, h.parentIndex(k))
		k = h.parentIndex(k)
	}
}

func (h *MaxHeap) ShiftDown(i int) {
	for h.leftChildIndex(i) < h.data.GetSize() {
		leftChildIndex := h.leftChildIndex(i)
		rightChildIndex := leftChildIndex + 1
		// 获取左右孩子节点的较大值的索引
		if rightChildIndex < h.data.GetSize() && utils.Compare(h.data.Get(rightChildIndex), h.data.Get(leftChildIndex)) > 0 {
			leftChildIndex++
		}

		// 如果比左右孩子节点的较大值都大，则已经到合适的节点位置
		if utils.Compare(h.data.Get(i), h.data.Get(leftChildIndex)) > 0 {
			break
		}

		// 交换i和较大值索引位置，并更新i
		h.data.Swap(i, leftChildIndex)
		i = leftChildIndex
	}
}
