package heap

type Heap interface {
	ShiftDown(int)
	ShiftUp(k int)
	ExtractMax() interface{}
	Replace(e interface{}) interface{}
	Add(e interface{})
	FindMax() interface{}
	IsEmpty() bool
	Size() int
}
