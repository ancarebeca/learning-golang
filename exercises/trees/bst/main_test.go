package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertNodeAndTraversals(t *testing.T) {
	bst := createTree([]int{50, 30, 20, 40, 70, 60, 80})

	assert.Equal(t, []int{50, 30, 70, 20, 40, 60, 80}, breadthFirstSearch(bst.root))

	assert.Equal(t, []int{50, 30, 20, 40, 70, 60, 80}, dfsPreOrder(bst.root))
	assert.Equal(t, []int{50, 30, 20, 40, 70, 60, 80}, dfsPreOrderRec(bst.root))

	assert.Equal(t, []int{20, 30, 40, 50, 60, 70, 80}, dfsInorder(bst.root))
	assert.Equal(t, []int{20, 30, 40, 50, 60, 70, 80}, dfsInorderRec(bst.root))

	assert.Equal(t, []int{20, 40, 30, 60, 80, 70, 50}, dfsPostOrder(bst.root))
	assert.Equal(t, []int{20, 40, 30, 60, 80, 70, 50}, dfsPostOrderRec(bst.root))
}

func TestHeight(t *testing.T) {
	bst := createTree([]int{50, 30, 20, 40, 70, 60, 80})
	assert.Equal(t, 2, getHeight(bst.root))
}

func TestSearch(t *testing.T) {
	bst := createTree([]int{50, 30, 20, 40, 70, 60, 80})

	node := find(bst.root, 50)
	assert.Equal(t, 50, node.Val)

	node = find(bst.root, 30)
	assert.Equal(t, 30, node.Val)

	node = find(bst.root, 80)
	assert.Equal(t, 80, node.Val)

	node = find(bst.root, 90)
	assert.Nil(t, node)
}

func TestDelete(t *testing.T) {
	bst1 := createTree([]int{50, 30, 20, 40, 70, 60, 80})
	delete(bst1.root, 50)
	assert.Equal(t, []int{60, 30, 70, 20, 40, 80}, breadthFirstSearch(bst1.root))

	bst2 := createTree([]int{50, 30, 20, 40, 70, 60, 80})
	delete(bst2.root, 20)
	assert.Equal(t, []int{50, 30, 70, 40, 60, 80}, breadthFirstSearch(bst2.root))

	bst3 := createTree([]int{50, 30, 20, 40, 70, 60, 55, 80})
	delete(bst3.root, 60)
	assert.Equal(t, []int{50, 30, 70, 20, 40, 55, 80}, breadthFirstSearch(bst3.root))
}

func TestCheckBalanced(t *testing.T) {
	bst := createTree([]int{50, 30, 80, 70, 90, 60, 75})
	assert.False(t, isBalanced(bst.root))

	bst2 := createTree([]int{29, 16, 36, 10, 17, 34, 35, 6, 11, 20, 22})
	assert.False(t, isBalanced(bst2.root))

	bst3 := createTree([]int{50, 30, 20, 40, 70, 60, 80})
	assert.True(t, isBalanced(bst3.root))

	assert.False(t, isBalancedOptimised(bst2.root))
	assert.True(t, isBalancedOptimised(bst3.root))
}

func TestBalanceBST(t *testing.T) {
	bst := createTree([]int{50, 30, 80, 70, 90, 60, 75})
	balancesBST := balanceBST(bst.root)
	assert.True(t, isBalancedOptimised(balancesBST))

	bst2 := createTree([]int{29, 16, 36, 10, 17, 34, 35, 6, 11, 20, 22})
	balancesBST2 := balanceBST(bst2.root)
	assert.True(t, isBalancedOptimised(balancesBST2))

}

func createTree(values []int) BST {
	bst := BST{}
	for _, v := range values {
		bst.insert(v)
	}
	return bst
}
