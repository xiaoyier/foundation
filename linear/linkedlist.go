package linear

import "fmt"

type LinkNode struct {
	val interface{}
	next *LinkNode
}

type LinkedList struct {
	size int
	first *LinkNode
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList) Add(element interface{}) {
	ll.Insert(ll.size, element)
}

func (ll *LinkedList) Insert(index int, element interface{}) (ok bool) {
	if err := ll.rangeCheckAdd(index); err != nil {
		return false
	}
	if index == 0 {
		ll.first =  &LinkNode{
			val: element,
			next: ll.first,
		}
	} else {
		preNode, err := ll.nodeOfIndex(index-1)
		if err != nil {
			return false
		}
		preNode.next = &LinkNode{
			val: element,
			next: preNode.next,
		}
	}
	ll.size++
	return true
}

func (ll *LinkedList) RemoveAt(index int) (interface{}, error) {

	if err := ll.rangeCheck(index); err != nil {
		return nil, err
	}

	old := ll.first
	if index == 0 {
		ll.first = ll.first.next
	} else {
		preNode, err := ll.nodeOfIndex(index - 1)
		if err != nil {
			return nil, err
		}
		old = preNode.next
		preNode.next = old.next
	}

	ll.size--
	return old, nil
}

func (ll *LinkedList) RemoveFirst() (interface{}, error) {
	return ll.RemoveAt(0)
}

func (ll *LinkedList) RemoveLast() (interface{}, error) {
	return ll.RemoveAt(ll.size-1)
}

func (ll *LinkedList) RemoveAll() {
	ll.first = nil
	ll.size = 0
}

func (ll *LinkedList) Set(index int, element interface{}) error {
	node, err := ll.nodeOfIndex(index)
	if err != nil {
		return err
	}
	node.val = element
	return nil
}

func (ll *LinkedList) IndexOf(element interface{}) int {
	if element == nil {
		node := ll.first
		for i := 0; i < ll.size; i++ {
			if node == nil {
				return i
			}
			node = node.next
		}
	} else {
		node := ll.first
		for i := 0; i < ll.size; i++ {
			if element == node.val {
				return i
			}
			node = node.next
		}
	}
	return -1
}

func (ll *LinkedList) Get(index int) (interface{}, bool ){
	node, err := ll.nodeOfIndex(index)
	return node, err == nil
}

func (ll *LinkedList) Contains(element interface{}) bool {
	return ll.IndexOf(element) != -1
}

func (ll *LinkedList) nodeOfIndex(index int) (*LinkNode, error) {
	if err := ll.rangeCheck(index); err != nil {
		return nil, err
	}

	node := ll.first
	for i := 1; i <= index; i++ {
		node = node.next
	}

	return node, nil
}

func (list *LinkedList) rangeCheck (index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}
	return nil
}

func (list *LinkedList) rangeCheckAdd(index int) error {
	if index < 0 || index > list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}

	return nil
}