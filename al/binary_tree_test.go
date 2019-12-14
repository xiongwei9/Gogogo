package al

import (
	"testing"
)

type stack struct {
	data []interface{}
	size int
}

func (s *stack) push(data interface{}) {
	s.data = append(s.data, data)
	s.size += 1
}
func (s *stack) pop() interface{} {
	if s.size <= 0 {
		return nil
	}
	s.size -= 1
	data := s.data[s.size]
	s.data = s.data[:s.size]
	return data
}
func (s *stack) last() interface{} {
	if s.size <= 0 {
		return nil
	}
	return s.data[s.size-1]
}
func (s *stack) isEmpty() bool {
	return s.size <= 0
}

func createTree() *BST {
	data := []int{6, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	//tree := createBST(data)
	var tree *BST
	for _, v := range data {
		tree = insert(tree, v)
	}
	return tree
}

func compareArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, val := range arr1 {
		if val != arr2[i] {
			return false
		}
	}
	return true
}

func TestSearchTree(t *testing.T) {
	tree := createTree()

	if isExisted(tree, 5) && !isExisted(tree, 8) {
		t.Error("isExisted error 1")
	}

	//tree = &BST{val: 8}

	tree = treeDelete(tree, 8)
	if isExisted(tree, 8) {
		t.Error("delete error 2")
	}
}

func TestTraversalTreeFront(t *testing.T) {
	tree := createTree()

	test := []int{6, 4, 1, 3, 2, 16, 9, 8, 7, 10, 14}
	result := make([]int, 0, len(test))

	var mStack stack
	mStack.push(tree)
	for !mStack.isEmpty() {
		item := mStack.pop()
		node := item.(*BST)

		// 处理结点
		result = append(result, node.val)

		if node.right != nil {
			mStack.push(node.right)
		}
		if node.left != nil {
			mStack.push(node.left)
		}
	}

	if !compareArray(result, test) {
		t.Fatalf("arrays not equal:\n- %v\n- %v", test, result)
	}
}

func TestTraversalTreeMiddle(t *testing.T) {
	tree := createTree()

	test := []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 14, 16}
	result := make([]int, 0, len(test))

	var mStack stack
	node := tree

	for node != nil || !mStack.isEmpty() {
		for node != nil {
			mStack.push(node)
			node = node.left
		}

		if !mStack.isEmpty() {
			item := mStack.pop()
			node = item.(*BST)

			// 处理结点
			result = append(result, node.val)

			//mStack.pop()
			node = node.right
		}
	}

	if !compareArray(result, test) {
		t.Fatalf("arrays not equal:\n- %v\n- %v", test, result)
	}
}
