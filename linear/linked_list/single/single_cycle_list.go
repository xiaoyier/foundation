package single

import "fmt"

// 单向循环链表节点
type SCLinkdNode struct {
	val interface{}
	next *SCLinkdNode
}

// 单向循环链表
type SCLinkedList struct {
	size int
	first *SCLinkdNode
}

func (ll *SCLinkedList) Size() int {
	return ll.size
}

func (ll *SCLinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *SCLinkedList) Add(element interface{}) {
	ll.Insert(ll.size, element)
}

func (ll *SCLinkedList) Insert(index int, element interface{}) (ok bool) {
	if err := ll.rangeCheckAdd(index); err != nil {
		return false
	}

	if index == 0 {
		node :=  &SCLinkdNode{
			val: element,
			next: ll.first,
		}
		if ll.size == 0 {
			node.next = node
		} else {
			last, err := ll.nodeOfIndex(ll.size - 1)
			if err != nil {
				return false
			}
			last.next  = node
		}
		ll.first = node
	} else {
		preNode, err := ll.nodeOfIndex(index-1)
		if err != nil {
			return false
		}
		preNode.next = &SCLinkdNode{
			val: element,
			next: preNode.next,
		}
	}
	ll.size++
	return true
}

func (ll *SCLinkedList) RemoveAt(index int) (interface{}, error) {

	if err := ll.rangeCheck(index); err != nil {
		return nil, err
	}

	remove := ll.first
	if index == 0 {

		last, err := ll.nodeOfIndex(index)
		if err != nil {
			return nil, err
		}
		//只有一个节点
		if last == last.next {
			ll.first = nil
		} else {
			ll.first = remove.next
			last.next = ll.first
		}
	} else {
		preNode, err := ll.nodeOfIndex(index - 1)
		if err != nil {
			return nil, err
		}

		remove = preNode.next
		preNode.next = remove.next
	}

	ll.size--
	return remove, nil
}

func (ll *SCLinkedList) RemoveFirst() (interface{}, error) {
	return ll.RemoveAt(0)
}

func (ll *SCLinkedList) RemoveLast() (interface{}, error) {
	return ll.RemoveAt(ll.size-1)
}

func (ll *SCLinkedList) RemoveAll() {
	ll.first = nil
	ll.size = 0
}

func (ll *SCLinkedList) Set(index int, element interface{}) error {
	node, err := ll.nodeOfIndex(index)
	if err != nil {
		return err
	}
	node.val = element
	return nil
}

func (ll *SCLinkedList) IndexOf(element interface{}) int {
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

func (ll *SCLinkedList) Get(index int) (interface{}, bool ){
	node, err := ll.nodeOfIndex(index)
	return node, err == nil
}

func (ll *SCLinkedList) Contains(element interface{}) bool {
	return ll.IndexOf(element) != -1
}

func (ll *SCLinkedList) nodeOfIndex(index int) (*SCLinkdNode, error) {
	if err := ll.rangeCheck(index); err != nil {
		return nil, err
	}

	node := ll.first
	for i := 1; i <= index; i++ {
		node = node.next
	}

	return node, nil
}

func (list *SCLinkedList) rangeCheck (index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}
	return nil
}

func (list *SCLinkedList) rangeCheckAdd(index int) error {
	if index < 0 || index > list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}

	return nil
}