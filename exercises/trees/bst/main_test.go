package main

import (
	"fmt"
	"testing"
)

func TestInsertNode(t *testing.T) {
	bst := BST{}
	bst.insert(50)
	bst.insert(30)
	bst.insert(20)
	bst.insert(40)
	bst.insert(70)
	bst.insert(60)
	bst.insert(80)
	bst.bfsTranversal()
	fmt.Println("---")
	bst.bfsTranversalRec()

}
