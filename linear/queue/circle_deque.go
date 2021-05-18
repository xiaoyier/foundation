package queue


// 双端循环队列
type CircleDeque struct {
	head int
	size int
	elements []interface{}
}

func NewCircleDeque() *CircleDeque {
	return &CircleDeque{
		elements: make([]interface{}, DefauleQueueLength),
	}
}

func (q *CircleDeque) Size() int {
	return q.size
}

func (q *CircleDeque) IsEmpty() bool {
	return q.size == 0
}

func (q *CircleDeque) EnQueueFront(element interface{}) {
	// 扩容
	q.growCapacity(q.size + 1)

	var head int
	if q.head == 0 {
		head = len(q.elements) - 1
	} else {
		head = q.realIndex(q.head - 1)
	}

	q.elements[head] = element
	q.head = head
	q.size++
}

func (q *CircleDeque) EnQueueRear(element interface{}) {
	// 扩容
	q.growCapacity(q.size + 1)

	q.elements[q.realIndex(q.size)] = element
	q.size++
}

func (q *CircleDeque) DeQueueFront() interface{} {
	if q.size == 0 {
		return nil
	}

	element := q.elements[q.head]
	q.elements[q.head] = nil
	q.head = q.realIndex(q.head+1)
	q.size--
	return element
}

func (q *CircleDeque) DequeueRear() interface{} {
	if q.size == 0 {
		return nil
	}

	rear := q.realIndex(q.size-1)
	element := q.elements[q.realIndex(q.size-1)]
	q.elements[rear] = nil
	q.size--
	return element
}

func (q *CircleDeque) Front() (interface{}, bool) {
	if q.size == 0 {
		return nil, false
	}
	return q.elements[q.head], true
}

func (q *CircleDeque) Rear() (interface{}, bool) {
	if q.size == 0 {
		return nil, false
	}
	return q.elements[q.realIndex(q.size-1)], true
}

func (q *CircleDeque) Clear() {
	for i := 0; i < q.size; i++ {
		q.elements[q.realIndex(i)] = nil
	}
	q.size = 0
	q.head = 0
}

func (q *CircleDeque) realIndex(index int) int {
	if q.head + index < len(q.elements) {
		return index
	}
	return len(q.elements) - (q.head + index)
	// 模运算更耗性能
	//return (q.head + index) % len(q.elements)
}

func (q *CircleDeque) growCapacity(capacity int) {

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
