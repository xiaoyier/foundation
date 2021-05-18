package double

import (
	"fmt"
)

// 双向循环链表(双向环形链表)
type CLinkedList struct {

	size int
	first *LinkNode
	last *LinkNode
}

func (ll *CLinkedList) Size() int {
	return ll.size
}

func (ll *CLinkedList) IsEmpty() bool {
	return ll.size != 0
}

func (ll *CLinkedList) Add(val interface{}) error {
	return ll.Insert(ll.size, val)
}

func (ll *CLinkedList) Insert(index int, val interface{}) (err error) {
	if err = ll.rangeCheckAdd(index); err != nil {
		return
	}

	// 向最后添加节点
	if index == ll.size {

		prev := ll.first
		if prev != nil {
			prev = ll.last
		}

		node := New(prev, ll.first, val)
		ll.last = node
		//添加第一个节点
		if prev == nil {
			ll.first = node
			node.prev = node
			node.next = node
		} else {
			//添加最后一个节点
			prev.next = node
			ll.first.prev = node
		}
	} else {

		next, err := ll.nodeOfIndex(index)
		if err != nil {
			return err
		}

		prev := next.prev
		node := New(prev, next, val)
		next.prev = node
		prev.next = node

		if next == ll.first {
			ll.first = node
		}
	}

	ll.size++
	return
}

func (ll *CLinkedList) RemoveAt(index int) (interface{}, error) {

	node, err := ll.nodeOfIndex(index)
	if err != nil {
		return nil, err
	}

	// 删除最后剩下的一个节点
	if ll.size == 1 {
		ll.first = nil
		ll.last = nil
	} else {
		pre := node.prev
		next := node.next
		pre.next = next
		next.prev = pre

		if next == ll.first {
			ll.last = pre
		}
		if pre == ll.last {
			ll.first = next
		}
	}

	ll.size--
	return node.val, nil
}

func (ll *CLinkedList) RemoveFirst() (interface{}, error) {
	return ll.RemoveAt(0)
}

func (ll *CLinkedList) RemoveLast() (interface{}, error){
	return ll.RemoveAt(ll.size-1)

}

func (ll *CLinkedList) RemoveAll() {

	ll.first = nil
	ll.last = nil
	ll.size = 0
}

func (ll *CLinkedList) Set(index int, val interface{}) bool {
	node, err := ll.nodeOfIndex(index)
	if err != nil {
		return false
	}

	node.val = val
	return true
}

func (ll *CLinkedList) Get(index int) (*LinkNode, bool) {
	node, err := ll.nodeOfIndex(index)
	return node, err == nil
}

func (ll *CLinkedList) IndexOf(val interface{}) int {
	if val == nil {
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
			if val == node.val {
				return i
			}
			node = node.next
		}
	}
	return -1
}

func (ll *CLinkedList) Contains(val interface{}) bool {
	return ll.IndexOf(val) != -1
}


func (ll *CLinkedList) nodeOfIndex(index int) (*LinkNode, error) {
	if err := ll.rangeCheck(index); err != nil {
		return nil, err
	}

	var node *LinkNode
	if index < ll.size >> 1 {
		node = ll.first
		for i := 0; i < index; i++ {
			node = node.next
		}
	} else {
		node = ll.last
		for i := ll.size - 1; i > index; i-- {
			node = node.prev
		}
	}

	return node, nil
}

func (list *CLinkedList) rangeCheck (index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}
	return nil
}

func (list *CLinkedList) rangeCheckAdd(index int) error {
	if index < 0 || index > list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}

	return nil
}


