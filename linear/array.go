package linear

import (
	"bytes"
	"fmt"
)

const (
	ArrayDefaultCapacity = 10
	ArrayItemNotFount = -1
)


type ArrayItemer interface {
	Equal(another ArrayItemer)
	Reset()
}


type ArrayList struct {
	size int
	elements []interface{}
}

func Init(capacity int) *ArrayList {

	if capacity < 10 {
		capacity = ArrayDefaultCapacity
	}

	return &ArrayList{
		elements: make([]interface{}, capacity),
	}
}

func (list *ArrayList) Size() int {
	return list.size
}

func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

func (list *ArrayList) Append(elements ...interface{}) {

	list.growCapacity(list.size + len(elements))
	for _, element := range elements {
		list.elements[list.size] = element
		list.size++
	}
}

func (list *ArrayList) Insert(index int, element interface{}) error {
	err := list.rangeCheck(index)
	if err != nil {
		return err
	}

	list.growCapacity(list.size+1)
	for i := list.size; i > index; i-- {
		list.elements[i] = list.elements[i-1]
	}

	list.elements[index] = element
	list.size++
	return nil
}

func (list *ArrayList) Remove(element interface{}) int {
	index := list.IndexOf(element)
	if index == ArrayItemNotFount {
		return index
	}

	_, err := list.RemoveAt(index)
	if err != nil {
		return ArrayItemNotFount
	}

	return index
}

func (list *ArrayList) RemoveAt(index int) (interface{}, error) {
	err := list.rangeCheck(index)
	if err != nil {
		return nil, err
	}

	element := list.elements[index]

	//0 1 2 3 4 5
	for i := index; i < list.size; i++ {
		list.elements[i] = list.elements[i+1]
	}

	list.size--
	list.elements[list.size] = nil
	return element, nil
}

func (list *ArrayList) RemoveFirst() {
	for i := 0; i < list.size; i++ {
		list.elements[i] = list.elements[i+1]
	}

	list.size--
	list.elements[list.size] = nil
}

func (list *ArrayList) RemoveLast() {
	list.size--
	list.elements[list.size] = nil
}

func (list *ArrayList) RemoveAll() {
	for index := range list.elements {
		list.elements[index] = nil
	}
	list.size = 0
}

func (list *ArrayList) Set(index int, element interface{}) error {
	err := list.rangeCheck(index)
	if err != nil {
		return err
	}

	list.elements[index] = element
	return nil
}

func (list *ArrayList) Contains(element interface{}) bool {
	return list.IndexOf(element) != ArrayItemNotFount
}

func (list *ArrayList) IndexOf(element interface{}) int {
	if element == nil {
		for index, ele := range list.elements {
			if ele == nil {return index}
		}
	} else {
		for index, ele := range list.elements {
			if element == ele {return index}
		}
	}

	return ArrayItemNotFount
}

func (list *ArrayList) Get(index int) (interface{}, bool) {
	err := list.rangeCheck(index)
	if err != nil {
		return nil, false
	}
	return list.elements[index], true
}

func (list *ArrayList) SubArray(indexs ...int) *ArrayList {
	err := list.rangeCheck(indexs[len(indexs)-1])
	if err != nil {
		return nil
	}

	newElement := make([]interface{}, list.size)
	for index, val := range indexs {
		newElement[index] = list.elements[val]
	}

	return &ArrayList{
		size: len(indexs),
		elements: newElement,
	}
}

func (list *ArrayList) rangeCheck (index int) error {
	if index < 0 || index >= list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}
	return nil
}

func (list *ArrayList) rangeCheckAdd(index int) error {
	if index < 0 || index > list.size {
		return fmt.Errorf("Out of bound, size: %d, index: %d ", list.size, index)
	}

	return nil
}

func (list *ArrayList) growCapacity(capacity int) {

	oldCapacity := len(list.elements)
	if oldCapacity >= capacity {
		return
	}

	newCapacity := capacity + capacity >> 1
	newElements := make([]interface{}, newCapacity)
	for index, ele := range list.elements {
		newElements[index] = ele
	}

	list.elements = newElements
}

func (list *ArrayList) String() string  {

	buffer := new(bytes.Buffer)
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", list.size, len(list.elements)))
	buffer.WriteString("[")
	for i := 0; i < list.size; i++ {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprint(list.elements[i]))
	}

	buffer.WriteString("]")
	return buffer.String()
}


