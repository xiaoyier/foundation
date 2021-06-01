package binary

import "math"

type AVLTreeNode struct {
	TreeNode
	height int
}

func (n *AVLTreeNode) isBlanced() bool {
	leftHeight, rightHeight := 0, 0
	if n.Left() != nil {
		leftHeight = n.Left().(*AVLTreeNode).height
	}
	if n.Right() != nil {
		rightHeight = n.Right().(*AVLTreeNode).height
	}
	return math.Abs(float64(leftHeight-rightHeight)) < 1
}

func (n *AVLTreeNode) updateHeight() {
	leftHeight, rightHeight := 0, 0
	if n.Left() != nil {
		leftHeight = n.Left().(*AVLTreeNode).height
	}
	if n.Right() != nil {
		rightHeight = n.Right().(*AVLTreeNode).height
	}
	n.height = max(leftHeight, rightHeight) + 1
}

func (n *AVLTreeNode) tallerChildNode() *AVLTreeNode {
	leftHeight, rightHeight := 0, 0
	if n.Left() != nil {
		leftHeight = n.Left().(*AVLTreeNode).height
	}
	if n.Right() != nil {
		rightHeight = n.Right().(*AVLTreeNode).height
	}
	if leftHeight > rightHeight {
		return n.Left().(*AVLTreeNode)
	}
	return n.Right().(*AVLTreeNode)
}


type AVLTree struct {
	BinarySearchTree
}

func NewAVLTree(comparator NodeComparator) *AVLTree {
	return  &AVLTree{
		NewBinarySearchTree(comparator),
	}
}

func (t *AVLTree) AfterAdd(node TreeNode) {
	for node != nil {
		if t.isBalanced(node) {
			t.updateHeight(node)
		} else {
			t.reblance(node)
			break
		}
		node = node.Parent()
	}
}

func (t *AVLTree) AfterRemove(node TreeNode) {

	node = node.Parent()
	for  node != nil {
		if t.isBalanced(node) {
			t.updateHeight(node)
		} else {
			t.reblance(node)
		}
		node = node.Parent()
	}
}

func (t *AVLTree) reblance(node TreeNode) {
	g := node.(*AVLTreeNode)
	p := g.tallerChildNode()
	n := p.tallerChildNode()
	if p == g.Left() {
		if p.Left() == n { // LL, 右旋
			t.rotateRight(g,p,n)
		}else { // LR 右旋，左旋
			t.rotateRight(g,p,n)
			t.rotateLeft(g,p,n)
		}
	} else {
		if p.Left() == n { // RL 左旋，右旋
			t.rotateLeft(g,p,n)
			t.rotateRight(g,p,n)
		}else { // RR 左旋
			t.rotateLeft(g,p,n)
		}
	}
}

func (t *AVLTree) rotateLeft(g,p,n *AVLTreeNode) {
	g.SetRight(p.Left())
	p.SetLeft(g)

	bst := t.BinarySearchTree.(*binarySearchTree)
	if g.Parent() == nil {
		t.r
	}


}

func (t *AVLTree) rotateRight(g,p,n *AVLTreeNode) {
	g.SetLeft(p.Right())
	p.SetRight(g)
}

func (t *AVLTree) updateHeight(node TreeNode) {
	node.(*AVLTreeNode).updateHeight()
}

func (t *AVLTree) isBalanced(node TreeNode) bool {
	return node.(*AVLTreeNode).isBlanced()
}







