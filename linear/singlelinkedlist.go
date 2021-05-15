package linear

import "fmt"

type SingleLinkNode struct {
	val interface{}
	next *SingleLinkNode
}

type SingleLinkedList struct {
	size int
	first *SingleLinkNode
}

func (ll *SingleLinkedList) Size() int {
	return ll.size
}

func (ll *SingleLinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *SingleLinkedList) Add(element interface{}) {
	ll.Insert(ll.size, element)
}

func (ll *SingleLinkedList) Insert(index int, element interface{}) (ok bool) {
	if err := ll.rangeCheckAdd(index); err != nil {
		return false
	}
	if index == 0 {
		ll.first =  &SingleLinkNode{
			val: element,
			next: ll.first,
		}
	} else {
		preNode, err := ll.nodeOfIndex(index-1)
		if err != nil {
			return false
		}
		preNode.next = &SingleLinkNode{
			val: element,
			next: preNode.next,
		}
	}
	ll.size++
	return true
}

func (ll *SingleLinkedList) RemoveAt(index int) (interface{}, error) {

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

func (ll *SingleLinkedList) RemoveFirst() (interface{}, error) {
	return ll.RemoveAt(0)
}

func (ll *SingleLinkedList) RemoveLast() (interface{}, error) {
	return ll.RemoveAt(ll.size-1)
}

func (ll *SingleLinkedList) RemoveAll() {
	ll.first = nil
	ll.size = 0
}

func (ll *SingleLinkedList) Set(index int, element interface{}) error {
	node, err := ll.nodeOfIndex(index)
	if err != nil {
		return err
	}
	node.val = element
	return nil
}

func (ll *SingleLinkedList) IndexOf(element interface{}) int {
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

func (ll *SingleLinkedList) Get(index int) (interface{}, bool ){
	node, err := ll.nodeOfIndex(index)
	return node, err == nil
}

func (ll *SingleLinkedList) Contains(element interface{}) bool {
	return ll.IndexOf(element) != -1
}

func (ll *SingleLinkedList) nodeOfIndex(index int) (*SingleLinkNode, error) {
	if err := ll.rangeCheck(index); err != nil {
		return nil, err
	}

	node := ll.first
	for i := 1; i <= index; i++ {
		node = node.next
	}

	return node, nil
}

func (list *SingleLinkedList) rangeCheck (index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}
	return nil
}

func (list *SingleLinkedList) rangeCheckAdd(index int) error {
	if index < 0 || index > list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}

	return nil
}