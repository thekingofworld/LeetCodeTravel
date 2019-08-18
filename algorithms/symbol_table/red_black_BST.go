package symbol_table

import "fmt"

const (
	RED   = true
	BLACK = false
)

type RedBlackBST struct {
	root *RedBlackNode
}

type RedBlackNode struct {
	key   int
	val   int
	left  *RedBlackNode
	right *RedBlackNode
	color bool
}

func (r *RedBlackBST) Get(key int) (val int, ok bool) {
	fNode := r.get(r.root, key)
	if fNode != nil {
		return fNode.val, true
	}
	return -1, false
}

func (r *RedBlackBST) get(node *RedBlackNode, key int) *RedBlackNode {
	if node == nil {
		return nil
	}
	if node.key > key {
		return r.get(node.left, key)
	} else if node.key < key {
		return r.get(node.right, key)
	} else {
		return node
	}
}

func (r *RedBlackBST) Put(key, val int) {
	r.root = r.put(r.root, key, val)
	r.root.color = BLACK
}

func (r *RedBlackBST) put(node *RedBlackNode, key, val int) *RedBlackNode {
	if node == nil {
		return &RedBlackNode{key: key, val: val, left: nil, right: nil, color: RED}
	}
	if node.key > key {
		node.left = r.put(node.left, key, val)
	} else if node.key < key {
		node.right = r.put(node.right, key, val)
	} else {
		node.val = val
	}
	if r.isRed(node.right) && !r.isRed(node.left) {
		node = r.rotateLeft(node)
	}
	if r.isRed(node.left) && r.isRed(node.left.left) {
		node = r.rotateRight(node)
	}
	if r.isRed(node.left) && r.isRed(node.right) {
		r.flipColors(node)
	}
	return node
}

func (r *RedBlackBST) rotateLeft(node *RedBlackNode) *RedBlackNode {
	x := node.right
	node.right = x.left
	x.left = node
	x.color = node.color
	node.color = RED
	return x
}

func (r *RedBlackBST) rotateRight(node *RedBlackNode) *RedBlackNode {
	x := node.left
	node.left = x.right
	x.right = node
	x.color = node.color
	node.color = RED
	return x
}

func (r *RedBlackBST) flipColors(node *RedBlackNode) {
	node.color = RED
	node.left.color = BLACK
	node.right.color = BLACK
}

func (r *RedBlackBST) isRed(node *RedBlackNode) bool {
	if node == nil {
		return false
	}
	return node.color == RED
}

func (r *RedBlackBST) Print() {
	r.print(r.root)
}

func (r *RedBlackBST) print(node *RedBlackNode) {
	if node == nil {
		return
	}
	r.print(node.left)
	fmt.Println(node.key)
	r.print(node.right)
}
