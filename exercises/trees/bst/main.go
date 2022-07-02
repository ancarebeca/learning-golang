package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

type BST struct {
	root *Node
}

func (bst *BST) insert(key int) {
	bst.root = insertRec(bst.root, key)
}

// Time O(n) where n is the number of nodes
// Space O(n) with 'n' being the depth of the tree since at any point of time maximum number of stack frames that could
// be present in memory is 'n'.
func insertRec(node *Node, value int) *Node {
	if node == nil {
		return &Node{
			Val: value,
		}
	}

	if value > node.Val {
		node.Right = insertRec(node.Right, value)
	}

	if value < node.Val {
		node.Left = insertRec(node.Left, value)
	}

	return node
}

//  Delete
// Time: O(n)
// Space: O(n)
func delete(root *Node, value int) *Node {

	// base case: the value is not found in the tree
	if root == nil {
		return root
	}

	// if the given value is less than the root node, recur for the left subtree
	if value < root.Val {
		root.Left = delete(root.Left, value)
	}

	// if the given value is bigger than the root node, recur for the right subtree
	if value > root.Val {
		root.Right = delete(root.Right, value)
	}

	// value found
	if root.Val == value {
		//Case 1: node to be deleted has no children (it is a leaf node)
		if root.Left == nil && root.Right == nil {
			return nil
		}
		//Case 2: Deleting a node with one child: remove the node and replace it with its child.
		if root.Left == nil && root.Right != nil {
			return root.Right
		}

		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		// Case3: Deleting a node with two children.
		tempNode := minValued(root.Right)             // Find min right sub-tree
		root.Val = tempNode.Val                       // Copy the value in targeted node
		root.Right = delete(root.Right, tempNode.Val) // Delete duplicate from right subtree
	}

	return root
}

func minValued(root *Node) *Node {
	temp := root
	for nil != temp && temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

// Time: O(n)
// Space: O(n)
func breadthFirstSearch(node *Node) []int {
	var output []int
	var queue []*Node
	queue = append(queue, node)

	for len(queue) > 0 {
		current := queue[0]
		output = append(output, current.Val)
		queue = queue[1:]

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return output
}

func (bst *BST) breadthFirstSearchRec() {
	var q queue
	q.add(bst.root)
	bst.searchRec(q)
}

func (bst *BST) searchRec(q queue) {

	if len(q.items) <= 0 {
		return
	}

	current := q.pop()
	fmt.Println(current.Val)

	if current.Left != nil {
		q.add(current.Left)
	}

	if current.Right != nil {
		q.add(current.Right)
	}

	bst.searchRec(q)
}

// ---- Depth-first searchRec

// Pre-order
func dfsPreOrder(node *Node) []int {
	var output []int
	stack := []*Node{}
	stack = append(stack, node)

	for len(stack) > 0 {

		current := stack[len(stack)-1]
		output = append(output, current.Val)
		stack = stack[:len(stack)-1]

		if current.Right != nil {
			stack = append(stack, current.Right)
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
		}

	}

	return output
}

func dfsPreOrderRec(node *Node) []int {
	if node == nil {
		return nil
	}

	var l, r, o []int
	if node.Left != nil {
		l = dfsPreOrderRec(node.Left)
	}

	if node.Right != nil {
		r = dfsPreOrderRec(node.Right)
	}
	o = append(o, node.Val)
	o = append(o, l...)
	o = append(o, r...)

	return o
}

// In-order
func dfsInorder(node *Node) []int {
	var output []int
	stack := []*Node{}

	current := node
	for len(stack) > 0 || current != nil {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
			continue
		}

		current = stack[len(stack)-1]
		output = append(output, current.Val)
		current = current.Right
		stack = stack[:len(stack)-1]
	}

	return output
}

func dfsInorderRec(node *Node) []int {
	var o []int
	if node == nil {
		return nil
	}

	left := dfsInorderRec(node.Left)
	right := dfsInorderRec(node.Right)
	o = append(o, left...)
	o = append(o, node.Val)
	o = append(o, right...)
	return o
}

// Post-oder
func dfsPostOrderRec(node *Node) []int {
	var o []int
	if node == nil {
		return nil
	}

	left := dfsPostOrderRec(node.Left)

	right := dfsPostOrderRec(node.Right)

	o = append(o, left...)
	o = append(o, right...)
	o = append(o, node.Val)

	return o
}

func dfsPostOrder(node *Node) []int {
	var o []int
	if node == nil {
		return nil
	}

	left := dfsPostOrderRec(node.Left)

	right := dfsPostOrderRec(node.Right)

	o = append(o, left...)
	o = append(o, right...)
	o = append(o, node.Val)

	return o
}

// Height
func getHeight(node *Node) int {
	var q []*Node
	q = append(q, node)
	height := 0

	for len(q) > 0 {
		// calculate the total number of nodes at the current level
		size := len(q)

		for size > 0 {
			current := q[0]
			q = q[1:]

			if current.Left != nil {
				q = append(q, current.Left)
			}

			if current.Right != nil {
				q = append(q, current.Right)
			}
			size--
		}
		// increment height by 1 for each level
		height++
	}

	return height
}

func getHeightRec(node *Node) int {
	if node == nil {
		return -1
	}

	left := getHeightRec(node.Left)
	right := getHeightRec(node.Right)

	return 1 + max(left, right)
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

// Find the node for the given node value
func find(root *Node, value int) *Node {

	if root == nil {
		return nil
	}

	if root.Val == value {
		return root
	}

	if value < root.Val {
		return find(root.Left, value)
	}

	if value > root.Val {
		return find(root.Right, value)
	}

	return nil
}

// Check if BST is a balanced tree
func isBalanced(root *Node) bool {
	// To check if a tree is height-balanced, get the height of left and right subtrees. Return true if difference between
	// heights is not more than 1 and left and right subtrees are balanced, otherwise return false.
	if root == nil {
		return true
	}

	lh := getHeightRec(root.Left)
	rh := getHeightRec(root.Right)

	if math.Abs(float64(lh)-float64(rh)) > 1 {
		return false
	}

	if isBalanced(root.Left) && isBalanced(root.Right) == true {
		return true

	}

	// if we reach here means tree is not
	// height-balanced tree
	return false
}

func isBalancedOptimised(root *Node) bool {
	return isBalancedHelper(root) != -1
}

func isBalancedHelper(root *Node) int {
	if root == nil {
		return 0
	}

	lb := isBalancedHelper(root.Left)
	if lb < 0 {
		return -1
	}

	rb := isBalancedHelper(root.Right)
	if rb < 0 {
		return -1
	}

	if math.Abs(float64(lb)-float64(rb)) > 1 {
		return -1
	}

	return max(lb, rb) + 1
}

// Convert a normal BST to Balanced BST
// 1) Traverse given BST in inorder and store result in an array.
// 2) Build a balanced BST from the above created sorted array:
// 		2.1) Get the Middle of the array and make it root.
// 		2.2) Recursively do same for left half and right half.
//       a) Get the middle of left half and make it left child of the root
//          created in step 2.1.
//       b) Get the middle of right half and make it right child of the
//          root created in step 2.1
func balanceBST(root *Node) *Node {
	values := dfsInorderRec(root)
	return sortedArrayToBST(values, 0, len(values)-1)
}

func sortedArrayToBST(arr []int, start, end int) *Node {
	/* Base Case */
	if start > end {
		return nil
	}

	/* Get the middle element and make it root */
	mid := (start + end) / 2

	node := Node{Val: arr[mid]}

	/* Recursively construct the left subtree and make it
	   left child of root */
	node.Left = sortedArrayToBST(arr, start, mid-1)

	/* Recursively construct the right subtree and make it
	   right child of root */
	node.Right = sortedArrayToBST(arr, mid+1, end)

	return &node
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
