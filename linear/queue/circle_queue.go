package queue


// 循环队列
type CircleQueue struct {

	head int
	size int
	elements []interface{}
}

const DefauleQueueLength = 10

func NewCircleQueue() *CircleQueue {
	return &CircleQueue{
		elements: make([]interface{}, DefauleQueueLength),
	}
}

func (q *CircleQueue) Size() int {
	return q.size
}

func (q *CircleQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircleQueue) Enqueue(element interface{}) {
	// 扩容
	q.growCapacity(q.size + 1)

	q.elements[q.realIndex(q.size)] = element
	q.size++
}

func (q *CircleQueue) Dequeue() interface{} {
	if q.size == 0 {
		return nil
	}

	element := q.elements[q.head]
	q.elements[q.head] = nil
	q.head = q.realIndex(q.head+1)
	q.size--
	return element
}

func (q *CircleQueue) Front() (interface{}, bool) {
	if q.size == 0 {
		return nil, false
	}
	return q.elements[q.head], true
}

func (q *CircleQueue) Clear() {
	for i := 0; i < q.size; i++ {
		q.elements[q.realIndex(i)] = nil
	}
	q.size = 0
	q.head = 0
}

func (q *CircleQueue) realIndex(index int) int {
	if q.head + index < len(q.elements) {
		return index
	}
	return len(q.elements) - (q.head + index)
	// 模运算更耗性能
	//return (q.head + index) % len(q.elements)
}

func (q *CircleQueue) growCapacity(capacity int) {

	oldCapacity := len(q.elements)
	if oldCapacity >= capacity {
		return
	}

	newCapacity := capacity + capacity >> 1
	newElements := make([]interface{}, newCapacity)
	for index := range q.elements {
		realIndex := q.realIndex(index)
		newElements[index] = q.elements[realIndex]
	}

	q.elements = newElements
}



