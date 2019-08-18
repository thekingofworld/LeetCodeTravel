package symbol_table

import "fmt"

type BST struct {
	root *Node
}

type Node struct {
	key   int
	val   int
	left  *Node
	right *Node
}

func (b *BST) Get(key int) (val int, ok bool) {
	fNode := b.get(b.root, key)
	if fNode != nil {
		return fNode.val, true
	}
	return -1, false
}

func (b *BST) get(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key > key {
		return b.get(node.left, key)
	} else if node.key < key {
		return b.get(node.right, key)
	} else {
		return node
	}
}

func (b *BST) Put(key int, val int) {
	b.root = b.put(b.root, key, val)
}

func (b *BST) put(node *Node, key int, val int) *Node {
	if node == nil {
		return &Node{key: key, val: val, left: nil, right: nil}
	}
	if node.key > key {
		node.left = b.put(node.left, key, val)
	} else if node.key < key {
		node.right = b.put(node.right, key, val)
	} else {
		node.val = val
	}
	return node
}

func (b *BST) Delete(key int) {
	b.root = b.delete(b.root, key)
}

func (b *BST) delete(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key > key {
		node.left = b.delete(node.left, key)
	} else if node.key < key {
		node.right = b.delete(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		rightMinNode := b.min(node.right)
		rightMinNode.left = node.left
		rightMinNode.right = b.deleteMin(node.right)
		return rightMinNode
	}
	return node
}

func (b *BST) Min() (key int, ok bool) {
	node := b.min(b.root)
	if node == nil {
		return -1, false
	}
	return node.key, true
}

func (b *BST) min(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left != nil {
		return b.min(node.left)
	}
	return node
}

func (b *BST) Max() (key int, ok bool) {
	node := b.max(b.root)
	if node == nil {
		return -1, false
	}
	return node.key, true
}

func (b *BST) max(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.right != nil {
		return b.max(node.right)
	}
	return node
}

func (b *BST) DeleteMin() {
	b.root = b.deleteMin(b.root)
}

func (b *BST) deleteMin(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left != nil {
		node.left = b.deleteMin(node.left)
		return node
	}
	return node.right
}

func (b *BST) Print() {
	b.print(b.root)
}

func (b *BST) print(node *Node) {
	if node == nil {
		return
	}
	b.print(node.left)
	fmt.Println(node.key)
	b.print(node.right)
}
