package queue_stack

// 单调队列
// 生成一个队首为最大值的单调队列
type MonotonicQueue struct {
	queue []int
}

func NewMonoQueue() *MonotonicQueue {
	return &MonotonicQueue{}
}

func (mq *MonotonicQueue) Push(v int) {
	for len(mq.queue) > 0 && mq.queue[len(mq.queue)-1] < v {
		mq.queue = mq.queue[:len(mq.queue)-1]
	}
	mq.queue = append(mq.queue, v)
}

func (mq *MonotonicQueue) Pop(n int) {
	if len(mq.queue) > 0 && n == mq.queue[0] {
		mq.queue = mq.queue[1:len(mq.queue)]
	}
}

func (mq *MonotonicQueue) Max() int {
	return mq.queue[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	window := NewMonoQueue()
	var res []int
	if len(nums) < k {
		return res
	}
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			window.Push(nums[i])
		} else {
			window.Push(nums[i])
			res = append(res, window.Max())
			window.Pop(nums[i-k+1])
		}
	}
	return res
}
