package binary



type NodeComparator func (element1, element2 interface{}) int



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


func (t *BinarySearchTree) Add(element interface{}) {

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
	if element == nil {
		return
	}

	// 获取node
	node := t.nodeOf(element)
	if node == nil {
		return
	}

	// 度为2的节点
	if node.degree() == 2 {
		//找到其前驱或者后继节点
		predcessor := t.predcessor(node)
		// 覆盖内容
		node.element = predcessor.element
		// 删除 前驱节点
		node = predcessor
	}
	// 如果是叶子节点
	if node.isLeaf() {
		if node.parent == nil {
			t.root = nil
		} else if node == node.parent.left {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
	} else {
		//度为1的节点，找到替代自己的子节点
		child := node.left
		if node.right != nil {
			child = node.right
		}

		child.parent = node.parent
		if node.parent == nil {
			t.root = child
		} else if node == node.parent.left {
			node.parent.left = child
		} else {
			node.parent.right = child
		}
	}
}


func (t *BinarySearchTree) Contains(element interface{}) bool {
	return t.nodeOf(element) != nil
}

func (t *BinarySearchTree) nodeOf(element interface{}) *TreeNode {

	node := t.root
	for node != nil {
		cmp := t.comparator(node.element, element)
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



