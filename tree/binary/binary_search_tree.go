package binary

type TreeNode struct {
	element interface{}
	parent *TreeNode
	left *TreeNode
	right *TreeNode
}

func NewTreeNode(element interface{}, parent *TreeNode) *TreeNode {
	return &TreeNode{
		element: element,
		parent: parent,
	}
}

type Comparator func (element1, element2 interface{}) int

type BinarySearchTree struct {
	size int
	root *TreeNode
	comparator Comparator
}

func NewBinarySearchTree(comparator Comparator) *BinarySearchTree {
	return &BinarySearchTree{
		comparator: comparator,
	}
}

func (t *BinarySearchTree) Size() int {
	return t.size
}

func (t *BinarySearchTree) IsEmpty() bool {
	return t.size == 0
}

func (t *BinarySearchTree) Add(element interface{}) {
	// 添加根节点
	if t.root == nil {
		t.root = NewTreeNode(element, nil)
		t.size++
		return
	}

	node := t.root
	parent := node
	cmp := 0
	for node != nil {
		parent = node
		cmp = t.comparator(node.element, element)
		if cmp > 0 {
			node = node.left
		} else if cmp < 0 {
			node = node.right
		} else {
			node.element = element
			return
		}
	}

	newNode := NewTreeNode(element, parent)
	if cmp > 0 {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}


func (t *BinarySearchTree) Remove(element interface{}) {

}

func (t *BinarySearchTree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *BinarySearchTree) Contains(element interface{}) bool {
	return false
}

