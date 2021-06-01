package binary

type BinarySearchTreeNode interface {}


type NodeComparator interface {
	compare(element1, element2 interface{}) int
}

//type BinarySearchTree interface {
//	BinaryTree
//	Add(element interface{})
//	Remove(element interface{})
//	Contains(element interface{}) bool
//	AfterAdd(node TreeNode)
//	AfterRemove(node TreeNode)
//}

// 二叉搜索树
type BinarySearchTree struct {
	BinaryTree
	comparator NodeComparator
}

func NewBinarySearchTree(comparator NodeComparator) *BinarySearchTree {
	return &BinarySearchTree{
		comparator: comparator,
	}
}


func (t *BinarySearchTree) Add(element interface{}) *TreeNode {

	if element == nil {
		return nil
	}
	// 添加根节点
	if t.root == nil {
		t.root = NewTreeNode(element, nil)
		t.size++
		return nil
	}

	node := t.root
	parent := node
	cmp := 0
	for node != nil {
		parent = node
		cmp = t.comparator.compare(node.element, element)
		if cmp > 0 {
			node = node.left
		} else if cmp < 0 {
			node = node.right
		} else {
			node.element = element
			return nil
		}
	}

	newNode := NewTreeNode(element, parent)
	if cmp > 0 {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	return newNode
}


func (t *BinarySearchTree) Remove(element interface{}) *TreeNode {
	if element == nil {
		return nil
	}

	// 获取node
	node := t.nodeOf(element)
	if node == nil {
		return nil
	}

	// 度为2的节点
	if node.Degree() == 2 {
		//找到其前驱或者后继节点
		predcessor := t.predcessor(node)
		// 覆盖内容
		node.element = predcessor.element
		// 删除 前驱节点
		node = predcessor
	}

	child := node.left
	if node.right != nil {
		child = node.right
	}

	if child != nil {
		child.parent = node.parent
	}

	if node.parent == nil {
		t.root = child
	} else if node == node.parent.left {
		node.parent.left = child
	} else {
		node.parent.right = child
	}

	return node
}


func (t *BinarySearchTree) Contains(element interface{}) bool {
	return t.nodeOf(element) != nil
}

func (t *BinarySearchTree) nodeOf(element interface{}) *TreeNode {

	node := t.root
	for node != nil {
		cmp := t.comparator.compare(node.element, element)
		if cmp == 0 {
			return node
		} else if cmp > 0 {
			node = node.left
		} else {
			node = node.right
		}
	}

	return node
}




