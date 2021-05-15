package linear

import (
	"fmt"
)

type LinkNode struct {
	val interface{}
	prev *LinkNode
	next *LinkNode
}

func New(prev, next *LinkNode, val interface{}) *LinkNode {
	return &LinkNode{
		val: val,
		prev: prev,
		next: next,
	}
}

type LinkedList struct {

	size int
	first *LinkNode
	last *LinkNode
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.size != 0
}

func (ll *LinkedList) Add(val interface{}) error {
	return ll.Insert(ll.size, val)
}

func (ll *LinkedList) Insert(index int, val interface{}) error {
	if err := ll.rangeCheckAdd(index); err != nil {
		return err
	}

	// 向最后添加节点
	if index == ll.size {
		pre, err := ll.nodeOfIndex(index - 1)
		if err != nil {
			return err
		}
		node := New(pre, nil, val)

		ll.last = node
		if pre == nil {
			ll.first = node
		} else {
			pre.next = node
		}
	} else {
		next, err := ll.nodeOfIndex(index)
		if err != nil {
			return err
		}

		prev := next.prev
		node := New(next.prev, next, val)
		next.prev = node

		if prev == nil {
			ll.first = node
		} else {
			prev.next = node
		}
	}

	ll.size++
	return nil
}

func (ll *LinkedList) RemoveAt(index int) (interface{}, error) {

	node, err := ll.nodeOfIndex(index)
	if err != nil {
		return nil, err
	}

	pre := node.prev
	next := node.next

	// 首节点
	if pre == nil {
		ll.first = next
	} else {
		pre.next = next
	}

	//尾节点
	if next == nil {
		ll.last = pre
	} else {
		next.prev = pre
	}

	ll.size--
	return node.val, nil
}

func (ll *LinkedList) RemoveFirst() (interface{}, error) {
	return ll.RemoveAt(0)
}

func (ll *LinkedList) RemoveLast() (interface{}, error){
	return ll.RemoveAt(ll.size-1)

}

func (ll *LinkedList) RemoveAll() {

	ll.first = nil
	ll.last = nil
	ll.size = 0
}

func (ll *LinkedList) Set(index int, val interface{}) bool {
	 node, err := ll.nodeOfIndex(index)
	 if err != nil {
	 	return false
	 }

	 node.val = val
	 return true
}

func (ll *LinkedList) Get(index int) (*LinkNode, bool) {
	node, err := ll.nodeOfIndex(index)
	return node, err == nil
}

func (ll *LinkedList) IndexOf(val interface{}) int {
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

func (ll *LinkedList) Contains(val interface{}) bool {
	return ll.IndexOf(val) != -1
}


func (ll *LinkedList) nodeOfIndex(index int) (*LinkNode, error) {
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


