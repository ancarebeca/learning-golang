package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {

	assert.Equal(t, []int{1, 2, 8, 9, 12, 16}, displayNodes(reverse(createLinkedList([]int{1, 8, 2, 9, 16, 12}))))
	assert.Equal(t, []int{9, 12, 16, 4}, displayNodes(reverse(createLinkedList([]int{9, 4, 16, 12}))))
	assert.Equal(t, []int{1, 8, 2, 9, 12, 1}, displayNodes(reverse(createLinkedList([]int{1, 2, 8, 9, 12, 1}))))
	assert.Equal(t, []int{1, 9, 1}, displayNodes(reverse(createLinkedList([]int{1, 9, 1}))))

}

func createLinkedList(arr []int) *Node {
	head := &Node{}
	tempHead := head
	for _, v := range arr {
		if head == nil {
			head = &Node{data: v}
			tempHead = head
		} else {
			head.next = &Node{data: v}
			head = head.next
		}
	}
	head = tempHead.next
	return head

}
