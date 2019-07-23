package al_test

import (
	"testing"
)

func TestSearchTree(t *testing.T) {
	//var tree binaryTree
	data := []int{6, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	tree := &BST{val: data[0]}
	for _, v := range data[1:] {
		tree.insert(v)
	}

	if tree.isExisted(5) {
		t.Error("isExisted error 1")
	}
}

// Binary Tree
type BT interface {
	isExisted(val int) bool
	insert(val int)
	delete(val int) bool
}

// Binary Search Tree
type BST struct {
	val         int
	left, right *BST
}

// Red Black Tree
type RBT struct {
	BST
	mark bool
}

func (t *BST) isExisted(val int) bool {
	tmp := t
	for tmp != nil {
		if tmp.val == val {
			return true
		} else if tmp.val > val {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	return false
}

func (t *BST) insert(val int) {
	node := &BST{val: val}
	var prev *BST
	cur := t
	for cur != nil {
		prev = cur
		if cur.val > val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	if prev.val > val {
		prev.left = node
	} else {
		prev.right = node
	}
}
