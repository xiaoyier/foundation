package queue

import "foundation/linear/linked_list/double"

type Queue struct {
	linkedList *double.LinkedList
}

func (q *Queue) Size() int {
	return q.linkedList.Size()
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Enqueue(element interface{}) {
	_ = q.linkedList.Add(element)
}

func (q *Queue) Dequeue() interface{} {
	val, _ := q.linkedList.RemoveLast()
	return val
}

func (q *Queue) Front() (interface{}, bool) {
	return q.linkedList.Get(0)
}

func (q *Queue) Clear() {
	q.linkedList.RemoveAll()
}