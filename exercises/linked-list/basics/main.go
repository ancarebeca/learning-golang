package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*

 1  2  4
 1. 3. 4.

1.1 2 3. 4 4.
*/

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			cur.Next = list1
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	return head.Next
}

func main() {

}

func traverse(listNode *ListNode) {
	h := listNode

	for h != nil {
		fmt.Printf("%d -> ", h.Val)
		h = h.Next
	}
	fmt.Printf("NULL")
}
