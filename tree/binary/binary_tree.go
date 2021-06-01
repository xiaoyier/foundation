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


type WatchFunc func(element interface{}) bool

type nodeWatcher struct {
	watch WatchFunc
	stop bool
}


type TreeNode interface {
	IsLeaf() bool
	Degree() int
	Left() TreeNode
	Right() TreeNode
	Parent() TreeNode
	Element() interface{}
	SetLeft(node TreeNode)
	SetRight(node TreeNode)
	SetParent(node TreeNode)
	SetElement(element interface{})
}

// 二叉树节点
type treeNode struct {
	element interface{}
	parent TreeNode
	left TreeNode
	right TreeNode
}

func NewTreeNode(element interface{}, parent TreeNode) TreeNode {
	return &treeNode{
		element: element,
		parent: parent,
	}
}

func (n *treeNode) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *treeNode) Degree() int {
	degree := 0
	if n.left != nil {
		degree += 1
	}
	if n.right != nil {
		degree += 1
	}
	return degree
}

func (n *treeNode) Left() TreeNode {
	return n.left
}

func (n *treeNode) Right() TreeNode {
	return n.right
}

func (n *treeNode) Parent() TreeNode {
	return n.parent
}

func (n *treeNode) Element() interface{} {
	return n.element
}

func (n *treeNode) SetLeft(node TreeNode) {
	n.left = node
}
func (n *treeNode) SetRight(node TreeNode) {
	n.right = node
}

func (n *treeNode) SetParent(node TreeNode) {
	n.parent = node
}

func (n *treeNode) SetElement(element interface{}) {
	n.element = element
}

type BinaryTree interface {
	Size() int
	IsEmpty() bool
	Clear()
	IsCompleteTree() bool
	Height() int
}

type binaryTree struct {
	size int
	root TreeNode
}

func (t *binaryTree) Size() int {
	return t.size
}

func (t *binaryTree) IsEmpty() bool {
	return t.size == 0
}

func (t *binaryTree) Clear() {
	t.root = nil
	t.size = 0
}

// get predcessor node
func (t *binaryTree) predcessor(node TreeNode) TreeNode {
	// node.left.right.right.right.right
	if node == nil {
		return nil
	}

	// 拥有左子树
	if node.Left() != nil {
		node = node.Left()
		for node.Right() != nil {
			node = node.Right()
		}
		return node
	} else {
		// 无左子树，需要从父节点往上找
		for node.Parent() != nil && node.Parent().Left() == node {
			node = node.Parent()
		}

		// node.parent == nil(return nil) || node.parent.right == node (return node.parent)
		return node.Parent()
	}
}

//获取后继节点
func (t *binaryTree) successor(node TreeNode) TreeNode {
	// node.right.left.left.left.left
	if node == nil {
		return nil
	}
	if node.Right() != nil {
		node = node.Right()
		for node.Left() != nil {
			node = node.Left()
		}
		return node
	} else {
		// 无右子树，需要从父节点往上找
		for node.Parent() != nil && node.Parent().Right() == node {
			node = node.Parent()
		}

		// node.parent == nil(return nil) || node.parent.left == node (return node.parent)
		return node.Parent()
	}
}

func (t *binaryTree) Iterate(traversal BSTTraversalOrder, watch WatchFunc) {
	watcher := &nodeWatcher{
		watch: watch,
	}
	switch traversal {
	case BSTTraversalPreOrder:
		t.preOrder(watcher)
	case BSTTraversalInOrder:
		t.inOrder(watcher)
	case BSTTraversalPostOrder:
		t.postOrder(watcher)
	case BSTTraversalLevelOrder:
		t.levelOrder(watcher)
	default:
		fmt.Println("BSTTraversalOrder type error, choice:[BSTTraversalPreOrder, BSTTraversalInOrder, BSTTraversalPostOrder, BSTTraversalLevelOrder]")
	}
}

func (t *binaryTree) preOrder(watcher *nodeWatcher) {
	t.preOrderWithNode(t.root, watcher)
}

func (t *binaryTree) preOrderWithNode(node TreeNode, watcher *nodeWatcher) {
	if node == nil || watcher.stop {
		return
	}

	watcher.stop = watcher.watch(node.Element())
	t.preOrderWithNode(node.Left(), watcher)
	t.preOrderWithNode(node.Right(), watcher)
}

func (t *binaryTree) inOrder(watcher *nodeWatcher) {
	t.inOrderWithNode(t.root, watcher)
}

func (t *binaryTree) inOrderWithNode(node TreeNode, watcher *nodeWatcher) {
	if node == nil || watcher.stop {
		return
	}

	t.inOrderWithNode(node.Left(), watcher)
	if watcher.stop {
		return
	}
	watcher.stop = watcher.watch(node.Element())
	t.inOrderWithNode(node.Right(), watcher)
}

func (t *binaryTree) postOrder(watcher *nodeWatcher) {
	t.postOrderWithNode(t.root, watcher)
}

func (t *binaryTree) postOrderWithNode(node TreeNode, watcher *nodeWatcher) {
	if node == nil || watcher.stop {
		return
	}

	t.postOrderWithNode(node.Left(), watcher)
	t.postOrderWithNode(node.Right(), watcher)
	if watcher.stop {
		return
	}
	watcher.stop = watcher.watch(node.Element())
}

func (t *binaryTree) levelOrder(watcher *nodeWatcher) {
	q := queue.NewQueue()
	q.Enqueue(t.root)
	for !q.IsEmpty() {
		node := q.Dequeue().(TreeNode)
		watcher.stop = watcher.watch(node.Element())
		if watcher.stop {
			return
		}
		if node.Left() != nil {
			q.Enqueue(node.Left())
		}
		if node.Right() != nil {
			q.Enqueue(node.Right())
		}
	}
}

func (t *binaryTree) IsCompleteTree() bool {
	return IsCompleteTree(t.root)
}

// 是否是完全二叉树
// 完全二叉树的特征: 叶子节点只在最后2层，且度为1的节点最多只能有1个
func IsCompleteTree(root TreeNode) bool {

	q := queue.NewQueue()
	q.Enqueue(root)

	leaf := false
	for !q.IsEmpty() {
		node := q.Dequeue().(TreeNode)
		if node.Left() != nil {
			if leaf {
				return false
			}
			q.Enqueue(node.Left())
		} else if node.Right() != nil {
			return false
		}

		if node.Right() != nil {
			if leaf {
				return false
			}
			q.Enqueue(node.Right())
		} else {
			leaf = true
		}
	}

	return true
}

func (t *binaryTree) Height() int {
	//return Height(t.root)
	return Height2(t.root)
}

// 递归计算二叉树的高度
func Height(node TreeNode) int {
	if node == nil {
		return 0
	}

	// 递归计算
	return 1 + max(Height(node.Left()), Height(node.Right()))
}

// 遍历计算二叉树的高度
func Height2(node TreeNode) int {
	if node == nil {
		return 0
	}

	q := queue.NewQueue()
	q.Enqueue(node)
	height, levelSize := 0, 1
	for !q.IsEmpty() {

		node := q.Dequeue().(TreeNode)
		levelSize--

		if node.Left() != nil {
			q.Enqueue(node.Left())
		}

		if node.Right() != nil {
			q.Enqueue(node.Right())
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
