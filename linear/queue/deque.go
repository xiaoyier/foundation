package queue

import "foundation/linear/linked_list/double"

type Deque struct {
	linkedList *double.LinkedList
}

func (q *Deque) Size() int {
	return q.linkedList.Size()
}

func (q *Deque) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Deque) EnQueueFront(element interface{}) {
	_ = q.linkedList.Insert(0, element)
}

func (q *Deque) EnQueueRear(element interface{}) {
	_ = q.linkedList.Add(element)
}

func (q *Deque) DeQueueFront() interface{} {
	val, _ := q.linkedList.RemoveFirst()
	return val
}

func (q *Deque) DequeueRear() interface{} {
	val, _ := q.linkedList.RemoveLast()
	return val
}

func (q *Deque) Front() (interface{}, bool) {
	return q.linkedList.Get(0)
}

func (q *Deque) Rear() (interface{}, bool) {
	return q.linkedList.Get(q.linkedList.Size() - 1)
}

func (q *Deque) Clear() {
	q.linkedList.RemoveAll()
}