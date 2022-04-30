package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

type Query struct {
	u int
	c string
}

func find(node *Node, val int) *Node {
	if node.Val == val {
		return node
	}

	for i := 0; i < len(node.Children); i++ {
		found := find(node.Children[i], val)
		if found != nil {
			return found
		}
	}

	return nil
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

func countOfNodes(root *Node, queries []Query, s string) []int {
	dict := make(map[int]string)
	for i, c := range s {
		dict[i+1] = string(c)
	}

	var result []int
	for _, v := range queries {

		node := find(root, v.u)
		count := 0

		if node.Children == nil {
			continue
		}

		var q queue
		q.add(node)
		for len(q.items) > 0 {
			current := q.pop()

			if dict[current.Val] == v.c {
				count++
			}

			for i := 0; i < len(current.Children); i++ {
				q.add(current.Children[i])
			}
		}
		result = append(result, count)
	}

	return result
}

func printPretty(node *Node, indent string, last bool) {
	fmt.Printf("%s+- %v\n", indent, node.Val)
	if last {
		indent += "  "
	} else {
		fmt.Print("|-")
		indent += "| "
	}

	for i := 0; i < len(node.Children); i++ {
		printPretty(node.Children[i], indent, i == len(node.Children)-1)
	}
}
