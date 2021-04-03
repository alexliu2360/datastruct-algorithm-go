package queue_stack
//Go标准库里没有队列，可以用数组（切片）或链表来实现：
//
//使用数组切片；push就是append，pop就是调整切片长度，top就是返回最后一个元素
//使用标准库container/list包装
//自定义list，标准库的list是个双链表且将值定为interface{}类型，这里可以简化为单链表并确定数据类型为int

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func ConstructorMyStack() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var ret int
	if len(this.queue) > 0 {
		ret = this.queue[len(this.queue)-1]
		this.queue = this.queue[:len(this.queue)-1]
	}
	return ret
}

/** Get the top element. */
func (this *MyStack) Top() int {
	var ret int
	if len(this.queue) > 0 {
		ret = this.queue[len(this.queue)-1]
	}
	return ret
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
