package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func (n TreeNode) IsLeaf() bool {
	return n.Right == nil && n.Left == nil
}

func newNode(key int) *TreeNode {
	return &TreeNode{
		Val: key,
	}
}

type queue struct {
	items []*TreeNode
}

func (q *queue) add(item *TreeNode) {
	q.items = append(q.items, item)
}

func (q *queue) pop() *TreeNode {
	if len(q.items) <= 0 {
		return nil
	}

	el := q.items[0]
	q.items = q.items[1:]

	return el
}

//  Return the height of a tree, iterative
func height(root *TreeNode) int {
	var q queue
	q.add(root)
	height := 0

	for len(q.items) > 0 { //i
		// calculate the total number of nodes at the current level
		size := len(q.items)

		// process each node of the current level and enqueue their
		// non-empty Left and Right child
		for size > 0 { //j
			front := q.pop()

			if front.Left != nil {
				q.add(front.Left)
			}

			if front.Right != nil {
				q.add(front.Right)
			}

			size = size - 1
		}
		// increment heightRec by 1 for each level
		height = height + 1
	}
	return height
}

//  Return the heightRec of a tree, recursive
func heightRec(node *TreeNode) int {
	if node == nil || node.IsLeaf() {
		return 1
	}

	left := height(node.Left)
	right := height(node.Right)
	height := max(left, right)
	return 1 + height
}

// Recursive function to print right view of a binary tree.
func rightViewRec(node *TreeNode) []int {
	result := right(node)
	return result
}

func right(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}

	rs := rightViewRec(node.Right)
	ls := rightViewRec(node.Left)

	if len(ls) > len(rs) {
		rs = append(rs, ls[len(rs):]...)
	}

	res := make([]int, len(rs)+1)
	res[0] = node.Val
	copy(res[1:], rs)

	return res
}

// Iterative function to print right view of a binary tree.
func rightViewIterative(root *TreeNode) []int {
	var q queue
	q.add(root)
	height := 0
	var result []int

	if root == nil {
		return nil
	}

	for len(q.items) > 0 {
		size := len(q.items)
		result = append(result, q.items[0].Val)

		for size > 0 {
			front := q.pop()

			if front.Right != nil {
				q.add(front.Right)
			}

			if front.Left != nil {
				q.add(front.Left)
			}

			size = size - 1
		}
		height = height + 1
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
