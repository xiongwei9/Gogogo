package al

import (
	"testing"
)

func TestSearchTree(t *testing.T) {
	//var tree binaryTree
	data := []int{6, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	//tree := createBST(data)
	var tree *BST
	for _, v := range data {
		tree = insert(tree, v)
	}

	if isExisted(tree, 5) && !isExisted(tree, 8) {
		t.Error("isExisted error 1")
	}

	//tree = &BST{val: 8}

	tree = treeDelete(tree, 8)
	if isExisted(tree, 8) {
		t.Error("delete error 2")
	}
}
