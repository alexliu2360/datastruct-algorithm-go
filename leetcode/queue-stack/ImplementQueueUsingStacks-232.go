package queue_stack
type MyQueue struct {
	InputStack *stackArr
	OutputStack *stackArr
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		InputStack: &stackArr{
			arr: make([]int, 0),
			size: 0,
		},
		OutputStack: &stackArr{
			arr: make([]int, 0),
			size: 0,
		},
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.InputStack.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.OutputStack.IsEmpty(){
		for !this.InputStack.IsEmpty() {
			this.OutputStack.Push(this.InputStack.Pop())
		}
	}
	return this.OutputStack.Pop()

}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.OutputStack.IsEmpty(){
		for !this.InputStack.IsEmpty() {
			this.OutputStack.Push(this.InputStack.Pop())
		}
	}
	return this.OutputStack.Peek()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.OutputStack.IsEmpty() && this.InputStack.IsEmpty()
}


type stackArr struct {
	arr []int
	size int
}

func (s *stackArr)Pop()int{
	var ret int
	if s.size > 0{
		ret = s.arr[s.size-1]
		s.arr = s.arr[:s.size-1]
		s.size--
	}
	return ret
}

func (s *stackArr)Push(e int){
	s.arr = append(s.arr, e)
	s.size++
}

func (s *stackArr)Peek()int{
	var ret int
	if s.size > 0{
		ret = s.arr[s.size-1]
	}
	return ret
}

func (s *stackArr)IsEmpty()bool{
	return s.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */