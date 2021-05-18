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
	watcher func (element interface{}) bool
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
		 if node.left != nil {
		 	q.Enqueue(node.left)
		 }
		 if node.right != nil {
		 	q.Enqueue(node.right)
		 }
	}
}


