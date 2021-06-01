package binary



type NodeComparator interface {
	compare(element1, element2 interface{}) int
}

type BinarySearchTree interface {
	BinaryTree
	Add(element interface{})
	Remove(element interface{})
	Contains(element interface{}) bool
	AfterAdd(node TreeNode)
	AfterRemove(node TreeNode)
}

// 二叉搜索树
type binarySearchTree struct {
	BinaryTree
	comparator NodeComparator
}

func NewBinarySearchTree(comparator NodeComparator) BinarySearchTree {
	return &binarySearchTree{
		comparator: comparator,
	}
}


func (t *binarySearchTree) Add(element interface{}) {

	if element == nil {
		return
	}
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
		cmp = t.comparator.compare(node.Element(), element)
		if cmp > 0 {
			node = node.Left()
		} else if cmp < 0 {
			node = node.Right()
		} else {
			node.SetElement(element)
			return
		}
	}

	newNode := NewTreeNode(element, parent)
	if cmp > 0 {
		parent.SetLeft(newNode)
	} else {
		parent.SetParent(newNode)
	}

	t.AfterAdd(node)
}


func (t *binarySearchTree) Remove(element interface{}) {
	if element == nil {
		return
	}

	// 获取node
	node := t.nodeOf(element)
	if node == nil {
		return
	}

	// 度为2的节点
	if node.Degree() == 2 {
		//找到其前驱或者后继节点
		predcessor := t.predcessor(node)
		// 覆盖内容
		node.SetElement(predcessor.Element())
		// 删除 前驱节点
		node = predcessor
	}

	child := node.Left()
	if node.Right() != nil {
		child = node.Right()
	}

	if child != nil {
		child.SetParent(node.Parent())
	}

	if node.Parent() == nil {
		t.root = child
	} else if node == node.Parent().Left() {
		node.Parent().SetLeft(child)
	} else {
		node.Parent().SetRight(child)
	}
	t.AfterRemove(node)
}


func (t *binarySearchTree) Contains(element interface{}) bool {
	return t.nodeOf(element) != nil
}

func (t *binarySearchTree) nodeOf(element interface{}) TreeNode {

	node := t.root
	for node != nil {
		cmp := t.comparator.compare(node.Element(), element)
		if cmp == 0 {
			return node
		} else if cmp > 0 {
			node = node.Left()
		} else {
			node = node.Right()
		}
	}

	return node
}

func (t *binarySearchTree) AfterAdd(node TreeNode) {
	// implements by sub struct
}

func (t *binarySearchTree) AfterRemove(node TreeNode) {
	// implements by sub struct
}



