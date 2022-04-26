package main

import "fmt"

type Node struct {
	key   int
	right *Node
	left  *Node
}

type BST struct {
	root *Node
}

func (bst *BST) insert(key int) {
	bst.root = bst.insertRec(bst.root, key)
}

func (bst *BST) insertRec(node *Node, key int) *Node {

	if node == nil {
		return &Node{key: key}
	}

	if key < node.key {
		node.left = bst.insertRec(node.left, key)
	}
	if key > node.key {
		node.right = bst.insertRec(node.right, key)
	}

	return node
}

type queue struct {
	items []*Node
}

func (q *queue) add(item *Node) {
	q.items = append(q.items, item)
}

func (q *queue) pop() *Node {
	if len(q.items) <= 0 {
		return nil
	}

	el := q.items[0]
	q.items = q.items[1:]

	return el
}

// Breadth-first transversal
func (bst *BST) bfsTransversal() {
	var q queue
	q.add(bst.root)

	for len(q.items) > 0 {

		current := q.pop()
		fmt.Println(current.key)

		if current.left != nil {
			q.add(current.left)
		}

		if current.right != nil {
			q.add(current.right)
		}

	}
}

func (bst *BST) bfsTransversalRec() {
	var q queue
	q.add(bst.root)
	bst.bfsRec(q)
}

func (bst *BST) bfsRec(q queue) {

	if len(q.items) <= 0 {
		return
	}

	current := q.pop()
	fmt.Println(current.key)

	if current.left != nil {
		q.add(current.left)
	}

	if current.right != nil {
		q.add(current.right)
	}

	bst.bfsRec(q)
}

// Depth-first transversal

// Preorder <root><left><right>
func (bst *BST) preOrderTransversalRec() {
	preOrderRec(bst.root)
}

func preOrderRec(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.key)

	if node.left != nil {
		preOrderRec(node.left)
	}

	if node.right != nil {
		preOrderRec(node.right)
	}
}

type stack struct {
	items []*Node
}

func (s *stack) push(item *Node) {
	s.items = append(s.items, item)
}

func (s *stack) pop() *Node {
	if len(s.items) == 0 {
		return nil
	}

	// Top element
	top := len(s.items) - 1
	item := s.items[top]

	// Erase element (write zero value)
	s.items[top] = nil
	s.items = s.items[:top]

	return item
}

func (bst *BST) preOrderTransversal() {
	var s stack
	s.push(bst.root)

	for len(s.items) > 0 {
		n := s.pop()
		fmt.Println(n.key)

		if n.right != nil {
			s.push(n.right)
		}

		if n.left != nil {
			s.push(n.left)
		}
	}
}

// Inorder

// Postorder
