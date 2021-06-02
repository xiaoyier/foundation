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


// 二叉树节点
type TreeNode struct {
	element interface{}
	parent *TreeNode
	left *TreeNode
	right *TreeNode
	height int
}

func NewTreeNode(element interface{}, parent *TreeNode) *TreeNode {
	return &TreeNode{
		element: element,
		parent: parent,
	}
}

func (n *TreeNode) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *TreeNode) Degree() int {
	degree := 0
	if n.left != nil {
		degree += 1
	}
	if n.right != nil {
		degree += 1
	}
	return degree
}

type BinaryTree struct {
	size int
	root *TreeNode
}

func (t *BinaryTree) Size() int {
	return t.size
}

func (t *BinaryTree) IsEmpty() bool {
	return t.size == 0
}

func (t *BinaryTree) Clear() {
	t.root = nil
	t.size = 0
}

// get predcessor node
func (t *BinaryTree) predcessor(node *TreeNode) *TreeNode {
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
func (t *BinaryTree) successor(node *TreeNode) *TreeNode {
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

func (t *BinaryTree) Iterate(traversal BSTTraversalOrder, watch WatchFunc) {
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

func (t *BinaryTree) preOrder(watcher *nodeWatcher) {
	t.preOrderWithNode(t.root, watcher)
}

func (t *BinaryTree) preOrderWithNode(node *TreeNode, watcher *nodeWatcher) {
	if node == nil || watcher.stop {
		return
	}

	watcher.stop = watcher.watch(node.element)
	t.preOrderWithNode(node.left, watcher)
	t.preOrderWithNode(node.right, watcher)
}

func (t *BinaryTree) inOrder(watcher *nodeWatcher) {
	t.inOrderWithNode(t.root, watcher)
}

func (t *BinaryTree) inOrderWithNode(node *TreeNode, watcher *nodeWatcher) {
	if node == nil || watcher.stop {
		return
	}

	t.inOrderWithNode(node.left, watcher)
	if watcher.stop {
		return
	}
	watcher.stop = watcher.watch(node.element)
	t.inOrderWithNode(node.right, watcher)
}

func (t *BinaryTree) postOrder(watcher *nodeWatcher) {
	t.postOrderWithNode(t.root, watcher)
}

func (t *BinaryTree) postOrderWithNode(node *TreeNode, watcher *nodeWatcher) {
	if node == nil || watcher.stop {
		return
	}

	t.postOrderWithNode(node.left, watcher)
	t.postOrderWithNode(node.right, watcher)
	if watcher.stop {
		return
	}
	watcher.stop = watcher.watch(node.element)
}

func (t *BinaryTree) levelOrder(watcher *nodeWatcher) {
	q := queue.NewQueue()
	q.Enqueue(t.root)
	for !q.IsEmpty() {
		node := q.Dequeue().(TreeNode)
		watcher.stop = watcher.watch(node.element)
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

func (t *BinaryTree) IsCompleteTree() bool {
	return IsCompleteTree(t.root)
}

// 是否是完全二叉树
// 完全二叉树的特征: 叶子节点只在最后2层，且度为1的节点最多只能有1个
func IsCompleteTree(root *TreeNode) bool {

	q := queue.NewQueue()
	q.Enqueue(root)

	leaf := false
	for !q.IsEmpty() {
		node := q.Dequeue().(TreeNode)
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

func (t *BinaryTree) Height() int {
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

		node := q.Dequeue().(TreeNode)
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
