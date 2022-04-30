package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountOfNodes(t *testing.T) {
	root := Node{Val: 1}
	root.Children = append(root.Children, &Node{Val: 2})
	root.Children = append(root.Children, &Node{Val: 3})
	root.Children = append(root.Children, &Node{Val: 7})

	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 4})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 5})
	root.Children[1].Children = append(root.Children[1].Children, &Node{Val: 6})

	q := []Query{
		{1, "a"},
		{2, "b"},
		{3, "a"},
	}
	assert.Equal(t, []int{4, 1, 2}, countOfNodes(&root, q, "abaacab"))

	root2 := Node{Val: 1}
	root2.Children = append(root2.Children, &Node{Val: 2})
	root2.Children = append(root2.Children, &Node{Val: 3})

	q2 := []Query{
		{1, "a"},
	}
	assert.Equal(t, []int{2}, countOfNodes(&root, q2, "aba"))

	q3 := []Query{
		{2, "b"},
	}
	assert.Equal(t, []int{1}, countOfNodes(&root, q3, "aba"))

}
