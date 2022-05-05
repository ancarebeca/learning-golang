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

	head.next = curr
	return prev
}

func reverse(head *Node) *Node {

	tempHead := head
	prev := head
	curr := head

	for curr != nil {
		if curr.data%2 == 0 {
			v := reverseOdds(curr)
			prev.next = v
		}

		prev = curr
		curr = curr.next
	}

	return tempHead
}

func displayNodes(head *Node) []int {
	result := []int{}
	curr := head

	result = append(result, curr.data)
	for curr.next != nil {
		curr = curr.next
		result = append(result, curr.data)
	}

	return result
}
