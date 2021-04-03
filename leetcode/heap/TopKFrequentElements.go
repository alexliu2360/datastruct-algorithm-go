package heap

// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/description/
// 基于container/heap原生的组件
// 也可以自己实现一个堆

import (
	"code_brush/basic-data-structure/queue"
	"container/heap"
)

type freq struct {
	e         int    // 元素e
	frequency uint32 // e的出现频次
}

type freqMaxHeapDto []*freq

func (f freqMaxHeapDto) Len() int {
	return len(f)
}

func (f freqMaxHeapDto) Less(i, j int) bool {
	return f[i].frequency < f[j].frequency
}

func (f freqMaxHeapDto) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *freqMaxHeapDto) Push(x interface{}) {
	e := x.(*freq)
	*f = append(*f, e)
	heap.Fix(f, f.Len()-1)
}

func (f *freqMaxHeapDto) Pop() interface{} {
	e := (*f)[0]
	*f = (*f)[1 : len(*f)]
	heap.Fix(f, 0)
	return e
}

func (f *freqMaxHeapDto) Top() interface{} {
	e := (*f)[0]
	return e
}

func TopKFrequentElements(nums []int, k int) []int {
	numMaps := make(map[int]uint32)
	for _, num := range nums {
		numMaps[num]++
	}
	freqMaxHeap := make(freqMaxHeapDto, 0)

	heap.Init(&freqMaxHeap)
	for num, f := range numMaps {
		if freqMaxHeap.Len() < k {
			freqMaxHeap.Push(&freq{
				e:         num,
				frequency: f,
			})
		} else {
			if f > freqMaxHeap.Top().(*freq).frequency {
				freqMaxHeap.Pop()
				freqMaxHeap.Push(&freq{
					e:         num,
					frequency: f,
				})
			}
		}
	}
	var res []int
	for freqMaxHeap.Len() != 0 {
		res = append(res, freqMaxHeap.Pop().(*freq).e)
	}
	return res
}


func TopKFrequentElements2(nums []int, k int) []int {
	numMaps := make(map[int]uint32)
	for _, num := range nums {
		numMaps[num]++
	}
	p := queue.GetPriQueue()

	for num, f := range numMaps {
		if p.GetSize() < k {
			p.Enqueue(&freq{
				e:         num,
				frequency: f,
			})
		} else {
			if f > p.GetFront().(*freq).frequency {
				p.Dequeue()
				p.Enqueue(&freq{
					e:         num,
					frequency: f,
				})
			}
		}
	}

	var res []int
	for !p.IsEmpty() {
		res = append(res, p.Dequeue().(*freq).e)
	}
	return res
}
