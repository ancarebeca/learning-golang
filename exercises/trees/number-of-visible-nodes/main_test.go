package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeightIterative(t *testing.T) {
	root := newNode(8)
	root.Left = newNode(3)
	root.Right = newNode(10)
	root.Left.Left = newNode(1)
	root.Left.Right = newNode(6)
	root.Right.Right = newNode(14)
	root.Left.Right.Left = newNode(4)
	root.Left.Right.Right = newNode(7)
	root.Right.Right.Left = newNode(13)

	assert.Equal(t, 4, height(root))
}

func TestHeightRecursive(t *testing.T) {
	root := newNode(8)
	root.Left = newNode(3)
	root.Right = newNode(10)
	root.Left.Left = newNode(1)
	root.Left.Right = newNode(6)
	root.Right.Right = newNode(14)
	root.Left.Right.Left = newNode(4)
	root.Left.Right.Right = newNode(7)
	root.Right.Right.Left = newNode(13)

	assert.Equal(t, 4, heightRec(root))
}

func TestRightViewRec(t *testing.T) {
	root := newNode(8)
	root.Left = newNode(3)
	root.Right = newNode(10)
	root.Left.Left = newNode(1)
	root.Left.Right = newNode(6)
	root.Right.Right = newNode(14)
	root.Left.Right.Left = newNode(4)
	root.Left.Right.Right = newNode(7)
	root.Right.Right.Left = newNode(13)

	assert.Equal(t, []int{8, 10, 14, 13}, rightViewRec(root))
}

func TestRightViewIterative(t *testing.T) {
	tree1 := newNode(8)
	tree1.Left = newNode(3)
	tree1.Right = newNode(10)
	tree1.Left.Left = newNode(1)
	tree1.Left.Right = newNode(6)
	tree1.Right.Right = newNode(14)
	tree1.Left.Right.Left = newNode(4)
	tree1.Left.Right.Right = newNode(7)
	tree1.Right.Right.Left = newNode(13)

	assert.Equal(t, []int{8, 10, 14, 13}, rightViewIterative(tree1))

	tree2 := newNode(8)
	tree2.Left = newNode(3)
	tree2.Left.Left = newNode(1)

	assert.Equal(t, []int{8, 3, 1}, rightViewIterative(tree2))

	tree3 := newNode(8)
	tree3.Left = newNode(3)
	tree3.Right = newNode(10)
	tree3.Left.Left = newNode(1)
	tree3.Left.Right = newNode(6)
	tree3.Right.Right = newNode(14)
	tree3.Left.Right.Left = newNode(4)

	assert.Equal(t, []int{8, 10, 14, 4}, rightViewIterative(tree3))
}
