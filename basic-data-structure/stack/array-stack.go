package stack

import "code_brush/basic-data-structure/array"

type ArrayStack struct {
	array *array.Array
}

func New(capacity int) *ArrayStack {
	return &ArrayStack{
		array: array.New(capacity),
	}
}

func (as *ArrayStack) GetSize() int {
	return as.array.GetSize()
}

func (as *ArrayStack) IsEmpty() bool {
	return as.array.IsEmpty()
}

// 压入栈顶元素
func (as *ArrayStack) Push(element interface{}) {
	as.array.AddLast(element)
}

// 弹出栈顶元素
func (as *ArrayStack) Pop() interface{} {
	return as.array.RemoveLast()
}

// 查看栈顶元素
func (as *ArrayStack) Peek() interface{} {
	return as.array.GetLast()
}