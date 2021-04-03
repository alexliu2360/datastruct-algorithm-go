package set

import (
	linked_list "code_brush/basic-data-structure/linked-list"
)

type LinkedListSet struct {
	LLSet *linked_list.LinkedList
}

func NewLLSet() *LinkedListSet{
	return &LinkedListSet{
		LLSet: linked_list.New(),
	}
}

func (ll *LinkedListSet)Add(e interface{}){
	if !ll.LLSet.Contains(e) {
		ll.LLSet.AddFirst(e)
	}
}


func (ll *LinkedListSet) Remove(e interface{}) {
	ll.LLSet.RemoveElement(e)
}


func (ll *LinkedListSet) Contains(e interface{}) bool {
	return ll.LLSet.Contains(e)
}


func (ll *LinkedListSet) GetSize() int {
	return ll.LLSet.GetSize()
}

func (ll *LinkedListSet) IsEmpty() bool {
	return ll.LLSet.IsEmpty()
}
