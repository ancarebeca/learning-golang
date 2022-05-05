package main

type Node struct {
	data int
	next *Node
}

func reverseOdds(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil && curr.data%2 == 0 {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	head = prev
	return head
}

func reverse(head *Node) *Node {

	tempHead := &Node{data: 0, next: head}
	prev := tempHead
	curr := head

	for curr != nil {
		if curr.data%2 == 0 {
			prev.next = reverseOdds(curr)
		}

		prev = curr
		curr = curr.next
	}

	return tempHead.next
}
