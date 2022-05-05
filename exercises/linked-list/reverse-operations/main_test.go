package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {

	//fmt.Println(reverse(createLinkedList([]int{1, 8, 2, 9, 16, 12})))
	fmt.Println(reverse(createLinkedList([]int{9, 4, 16, 12})))

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
