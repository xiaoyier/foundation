package binary

import (
	"fmt"
	"foundation/linear/queue"
)

// 遍历方式
type BSTTraversalOrder int

const (
	BSTTraversalPreOrder  BSTTraversalOrder = iota	    	// 前序遍历
	BSTTraversalInOrder										// 中序遍历
	BSTTraversalPostOrder									// 后序遍历
	BSTTraversalLevelOrder									// 层序遍历
)

type NodeComparator func (element1, element2 interface{}) int
type NodeWatcher func (element interface{}) bool

type nodeWatch struct {
	stop bool
	watcher NodeWatcher
}

// 二叉树节点
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

func (n *TreeNode) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *TreeNode) degree() int {
	degree := 0
	if n.left != nil {
		degree += 1
	}
	if n.right != nil {
		degree += 1
	}
	return degree
}

// 二叉搜索树
type BinarySearchTree struct {
	size int
	root *TreeNode
	comparator NodeComparator
}

func NewBinarySearchTree(comparator NodeComparator) *BinarySearchTree {
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


func (t *BinarySearchTree) Clear() {
	t.root = nil
	t.size = 0
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

// 获取前驱节点
func (t *BinarySearchTree) predcessor(node *TreeNode) *TreeNode {
	// node.left.right.right.right.right
	if node == nil {
		return nil
	}

	// 拥有左子树
	if node.left != nil {
		node = node.left
		for node.right != nil {
			node = node.right
		}
		return node
	} else {
		// 无左子树，需要从父节点往上找
		for node.parent != nil && node.parent.left == node {
			node = node.parent
		}

		// node.parent == nil(return nil) || node.parent.right == node (return node.parent)
		return node.parent
	}
}

//获取后继节点
func (t *BinarySearchTree) successor(node *TreeNode) *TreeNode {
	// node.right.left.left.left.left
	if node == nil {
		return nil
	}
	if node.right != nil {
		node = node.right
		for node.left != nil {
			node = node.left
		}
		return node
	} else {
		// 无右子树，需要从父节点往上找
		for node.parent != nil && node.parent.right == node {
			node = node.parent
		}

		// node.parent == nil(return nil) || node.parent.left == node (return node.parent)
		return node.parent
	}
}

func (t *BinarySearchTree) Iterate(traversal BSTTraversalOrder, watcher NodeWatcher) {
	watch := &nodeWatch{
		watcher: watcher,
	}
	switch traversal {
	case BSTTraversalPreOrder:
		t.preOrder(watch)
	case BSTTraversalInOrder:
		t.inOrder(watch)
	case BSTTraversalPostOrder:
		t.postOrder(watch)
	case BSTTraversalLevelOrder:
		t.levelOrder(watch)
	default:
		fmt.Println("BSTTraversalOrder type error, choice:[BSTTraversalPreOrder, BSTTraversalInOrder, BSTTraversalPostOrder, BSTTraversalLevelOrder]")
	}
}

func (t *BinarySearchTree) preOrder(watcher *nodeWatch) {
	t.preOrderWithNode(t.root, watcher)
}

func (t *BinarySearchTree) preOrderWithNode(node *TreeNode, watcher *nodeWatch) {
	if node == nil || watcher.stop {
		return
	}

	watcher.stop = watcher.watcher(node.element)
	fmt.Println(node.element)
	t.preOrderWithNode(node.left, watcher)
	t.preOrderWithNode(node.right, watcher)
}

func (t *BinarySearchTree) inOrder(watcher *nodeWatch) {
	t.inOrderWithNode(t.root, watcher)
}

func (t *BinarySearchTree) inOrderWithNode(node *TreeNode, watcher *nodeWatch) {
	if node == nil || watcher.stop {
		return
	}

	t.inOrderWithNode(node.left, watcher)
	if watcher.stop {
		return
	}
	watcher.stop = watcher.watcher(node.element)
	t.inOrderWithNode(node.right, watcher)
}

func (t *BinarySearchTree) postOrder(watcher *nodeWatch) {
	t.postOrderWithNode(t.root, watcher)
}

func (t *BinarySearchTree) postOrderWithNode(node *TreeNode, watcher *nodeWatch) {
	if node == nil || watcher.stop {
		return
	}

	t.postOrderWithNode(node.left, watcher)
	t.postOrderWithNode(node.right, watcher)
	if watcher.stop {
		return
	}
	watcher.stop = watcher.watcher(node.element)
}

func (t *BinarySearchTree) levelOrder(watcher *nodeWatch) {
	q := queue.NewQueue()
	q.Enqueue(t.root)
	for !q.IsEmpty() {
		 node := q.Dequeue().(*TreeNode)
		 watcher.stop = watcher.watcher(node.element)
		 if watcher.stop {
			 return
		 }
		 if node.left != nil {
		 	q.Enqueue(node.left)
		 }
		 if node.right != nil {
		 	q.Enqueue(node.right)
		 }
	}
}

func (t *BinarySearchTree) IsCompleteTree() bool {
	return IsCompleteTree(t.root)
}

// 是否是完全二叉树
// 完全二叉树的特征: 叶子节点只在最后2层，且度为1的节点最多只能有1个
func IsCompleteTree(root *TreeNode) bool {

	q := queue.NewQueue()
	q.Enqueue(root)

	leaf := false
	for !q.IsEmpty() {
		node := q.Dequeue().(*TreeNode)
		if node.left != nil {
			if leaf {
				return false
			}
			q.Enqueue(node.left)
		} else if node.right != nil {
			return false
		}

		if node.right != nil {
			if leaf {
				return false
			}
			q.Enqueue(node.right)
		} else {
			leaf = true
		}
	}

	return true
}

func (t *BinarySearchTree) Height() int {
	//return Height(t.root)
	return Height2(t.root)
}

// 递归计算二叉树的高度
func Height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// 递归计算
	return 1 + max(Height(node.left), Height(node.right))
}

// 遍历计算二叉树的高度
func Height2(node *TreeNode) int {
	if node == nil {
		return 0
	}

	q := queue.NewQueue()
	q.Enqueue(node)
	height, levelSize := 0, 1
	for !q.IsEmpty() {

		node := q.Dequeue().(*TreeNode)
		levelSize--

		if node.left != nil {
			q.Enqueue(node.left)
		}

		if node.right != nil {
			q.Enqueue(node.right)
		}

		if levelSize == 0 {
			height++
			levelSize = q.Size()
		}
	}

	return height
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}


