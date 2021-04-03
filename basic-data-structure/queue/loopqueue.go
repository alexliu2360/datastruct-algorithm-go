package queue

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func NewLoopQueue(capacity int) *LoopQueue {
	lq := LoopQueue{
		data: make([]interface{}, capacity+1),
	}
	return &lq
}

func (lq *LoopQueue) GetCapacity() int {
	return len(lq.data) - 1
}

func (lq *LoopQueue) GetSize() int {
	return lq.size
}

func (lq *LoopQueue) IsEmpty() bool {
	return lq.front == lq.tail
}

func (lq *LoopQueue) Enqueue(e interface{}) {
	if (lq.tail+1)%len(lq.data) == lq.front {
		lq.resize(lq.GetCapacity() * 2)
	}

	lq.data[lq.tail] = e
	lq.tail = (lq.tail + 1) % len(lq.data)
	lq.size++
}

func (lq *LoopQueue) Dequeue() interface{} {
	if lq.IsEmpty() {
		panic("cannot dequeue from empty queue")
	}

	e := lq.data[lq.front]
	lq.data[lq.front] = nil
	// 循环队列需要执行求余运算
	lq.front = (lq.front + 1) % len(lq.data)
	lq.size--
	if lq.size == lq.GetCapacity()/4 && lq.size != 0 {
		lq.resize(lq.GetCapacity() / 2)
	}

	return e
}

func (lq *LoopQueue) resize(capacity int) {
	newData := make([]interface{}, capacity+1)
	for i := 0; i < lq.size; i++ {
		newData[i] = lq.data[(i+lq.front)%len(lq.data)]
	}

	lq.data = newData
	lq.front = 0
	lq.tail = lq.size
}
