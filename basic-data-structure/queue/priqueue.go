package queue

import "code_brush/basic-data-structure/heap"

// 基于heap的优先队列

type PriQueue struct {
	maxHeap *heap.MaxHeap
}

func GetPriQueue() *PriQueue {
	maxheap := heap.NewMaxHeap()
	return &PriQueue{
		maxHeap: maxheap,
	}
}

func (p *PriQueue) GetSize() int {
	return p.maxHeap.Size()
}

func (p *PriQueue) IsEmpty() bool {
	return p.maxHeap.IsEmpty()
}

func (p *PriQueue) Enqueue(e interface{}) {
	p.maxHeap.Add(e)
}

func (p *PriQueue) Dequeue() interface{} {
	return p.maxHeap.ExtractMax()
}

func (p *PriQueue) GetFront() interface{} {
	return p.maxHeap.FindMax()
}
