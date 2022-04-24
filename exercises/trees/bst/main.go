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

	if key < bst.root.key {
		node.left = bst.insertRec(node.left, key)
	}
	if key > bst.root.key {
		node.right = bst.insertRec(node.right, key)
	}

	return node
}

type queue struct {
	items []*Node
}

func (q *queue) add(item *Node) {
	items := q.items
	q.items = append(items, item)
}

func (q *queue) pop() *Node {
	if len(q.items) <= 0 {
		return nil
	}

	el := q.items[0]
	q.items = q.items[1:]

	return el
}

//
func (bst *BST) bfsTranversal() {
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

func (bst *BST) bfsTranversalRec() {
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
