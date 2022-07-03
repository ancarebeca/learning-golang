package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  6,
					Next: nil,
				},
			},
		},
	}

	l2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	h := MergeTwoLists(&l1, &l2)
	prev := h.Val

	l := h.Next
	count := 0
	for l != nil {
		count++
		assert.True(t, prev <= l.Val)
		prev = l.Val
		l = l.Next
	}
	assert.Equal(t, 8, count+1)

}

func TestMergeTwoListsA(t *testing.T) {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	l2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	h := MergeTwoLists(&l1, &l2)
	prev := h.Val

	l := h.Next
	count := 0
	for l != nil {
		count++
		assert.True(t, prev <= l.Val)
		prev = l.Val
		l = l.Next
	}

	assert.Equal(t, 6, count+1)
}

func TestMergeTwoListsB(t *testing.T) {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	l2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	h := MergeTwoLists(&l1, &l2)
	prev := h.Val

	l := h.Next
	count := 0
	for l != nil {
		count++
		assert.True(t, prev <= l.Val)
		prev = l.Val
		l = l.Next
	}

	assert.Equal(t, 6, count+1)
}

func TestMergeTwoListsC(t *testing.T) {
	var l1 *ListNode
	l2 := ListNode{
		Val: 1,
	}

	h := MergeTwoLists(l1, &l2)
	assert.Equal(t, 1, h.Val)
}

func TestMergeTwoListsD(t *testing.T) {
	var l1, l2 *ListNode

	h := MergeTwoLists(l1, l2)
	assert.Nil(t, h)
}

func TestMergeTwoListsE(t *testing.T) {
	var l2 *ListNode
	l1 := &ListNode{
		Val: 1,
	}

	h := MergeTwoLists(l1, l2)
	assert.Equal(t, 1, h.Val)
}

func TestMergeTwoListsF(t *testing.T) {
	l1 := ListNode{
		Val: 1,
	}

	l2 := ListNode{
		Val: 2,
	}

	h := MergeTwoLists(&l1, &l2)
	prev := h.Val

	l := h.Next
	count := 0
	for l != nil {
		count++
		assert.True(t, prev <= l.Val)
		prev = l.Val
		l = l.Next
	}

	assert.Equal(t, 2, count+1)
}
