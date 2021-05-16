package single

import (
	"fmt"
)

// 单向循环链表节点
type SLinkdNode struct {
	val interface{}
	next *SLinkdNode
}

// 单向循环链表
type SLinkedList struct {
	size int
	first *SLinkdNode
}

func (ll *SLinkedList) Size() int {
	return ll.size
}

func (ll *SLinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *SLinkedList) Add(element interface{}) {
	ll.Insert(ll.size, element)
}

func (ll *SLinkedList) Insert(index int, element interface{}) (ok bool) {
	if err := ll.rangeCheckAdd(index); err != nil {
		return false
	}
	if index == 0 {
		ll.first =  &SLinkdNode{
			val: element,
			next: ll.first,
		}
	} else {
		preNode, err := ll.nodeOfIndex(index-1)
		if err != nil {
			return false
		}
		preNode.next = &SLinkdNode{
			val: element,
			next: preNode.next,
		}
	}
	ll.size++
	return true
}

func (ll *SLinkedList) RemoveAt(index int) (interface{}, error) {

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

func (ll *SLinkedList) RemoveFirst() (interface{}, error) {
	return ll.RemoveAt(0)
}

func (ll *SLinkedList) RemoveLast() (interface{}, error) {
	return ll.RemoveAt(ll.size-1)
}

func (ll *SLinkedList) RemoveAll() {
	ll.first = nil
	ll.size = 0
}

func (ll *SLinkedList) Set(index int, element interface{}) error {
	node, err := ll.nodeOfIndex(index)
	if err != nil {
		return err
	}
	node.val = element
	return nil
}

func (ll *SLinkedList) IndexOf(element interface{}) int {
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

func (ll *SLinkedList) Get(index int) (interface{}, bool ){
	node, err := ll.nodeOfIndex(index)
	return node, err == nil
}

func (ll *SLinkedList) Contains(element interface{}) bool {
	return ll.IndexOf(element) != -1
}

func (ll *SLinkedList) nodeOfIndex(index int) (*SLinkdNode, error) {
	if err := ll.rangeCheck(index); err != nil {
		return nil, err
	}

	node := ll.first
	for i := 1; i <= index; i++ {
		node = node.next
	}

	return node, nil
}

func (list *SLinkedList) rangeCheck (index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}
	return nil
}

func (list *SLinkedList) rangeCheckAdd(index int) error {
	if index < 0 || index > list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}

	return nil
}