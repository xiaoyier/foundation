package binary

import (
	"math"
)


func (n *TreeNode) isBlanced() bool {
	leftHeight, rightHeight := 0, 0
	if n.left != nil {
		leftHeight = n.left.height
	}
	if n.right != nil {
		rightHeight = n.right.height
	}
	return math.Abs(float64(leftHeight-rightHeight)) < 1
}

func (n *TreeNode) updateHeight() {
	leftHeight, rightHeight := 0, 0
	if n.left != nil {
		leftHeight = n.left.height
	}
	if n.right != nil {
		rightHeight = n.right.height
	}
	n.height = max(leftHeight, rightHeight) + 1
}

func (n *TreeNode) tallerChildNode() *TreeNode {
	leftHeight, rightHeight := 0, 0
	if n.left != nil {
		leftHeight = n.left.height
	}
	if n.right != nil {
		rightHeight = n.right.height
	}
	if leftHeight > rightHeight {
		return n.left
	}
	return n.right
}


type AVLTree struct {
	*BinarySearchTree
}

func NewAVLTree(comparator NodeComparator) *AVLTree {
	return  &AVLTree{
		NewBinarySearchTree(comparator),
	}
}

func (t *AVLTree) Add(element interface{}) {
	node :=  t.BinarySearchTree.Add(element)
	for node != nil {
		if t.isBalanced(node) {
			t.updateHeight(node)
		} else {
			t.reblance(node)
			break
		}
		node = node.parent
	}
}

func (t *AVLTree) AfterRemove(node *TreeNode) {

	node = node.parent
	for  node != nil {
		if t.isBalanced(node) {
			t.updateHeight(node)
		} else {
			t.rebalance(node)
		}
		node = node.parent
	}
}

func (t *AVLTree) rebalance(node *TreeNode) {
	g := node
	p := g.tallerChildNode()
	n := p.tallerChildNode()
	if p == g.left {
		if n == p.left { // LL, 右旋
			t.rotateRight(g)
		}else { // LR 左旋，右旋
			t.rotateLeft(p)
			t.rotateRight(g)
		}
	} else {
		if n == p.left { // RL 右旋，左旋
			t.rotateRight(p)
			t.rotateLeft(g)
		}else { // RR 左旋
			t.rotateLeft(g)
		}
	}
}

func (t *AVLTree) rotateLeft(g *TreeNode) {
	p := g.right
	n := p.left
	g.right = n
	p.left = g

	if g.parent == nil {
		t.root = p
	} else if g.parent.left == g {
		g.parent.left = p
	} else {
		g.parent.right = p
	}

	p.parent = g.parent
	g.parent = p
	n.parent = g
}

func (t *AVLTree) rotateRight(g *TreeNode) {
	p := g.left
	n := p.right
	g.left = n
	p.right = g

	if g.parent == nil {
		t.root = p
	} else if g.parent.left == g {
		g.parent.left = p
	} else {
		g.parent.right = p
	}

	p.parent = g.parent
	g.parent = p
	n.parent = g
}

func (t *AVLTree) updateHeight(node *TreeNode) {
	node.updateHeight()
}

func (t *AVLTree) isBalanced(node *TreeNode) bool {
	return node.isBlanced()
}









