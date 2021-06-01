package binary

import (
	"fmt"
	"testing"
)


type compareStruct struct {
	compareFunc func(v1, v2 interface{}) int
}

func (c *compareStruct) compare(element1, element2 interface{}) int {
	return func(element1, element2 interface{}) int {
		v1 := element1.(int)
		v2 := element2.(int)
		return v1-v2
	}(element1, element2)
}

func TestBinarySearchTree(t *testing.T) {

	tree := NewBinarySearchTree(&compareStruct{})

	tree.Add(7)
	nums := [...]int{9,23,10,54,90,20,53,25,98,39,28,40,94,19,32}
	for _, num := range nums {
		tree.Add(num)
	}

	fmt.Println(tree)
}
