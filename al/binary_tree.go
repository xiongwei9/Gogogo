package al

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
	parent      *BST
}

// Red Black Tree
type RBT struct {
	BST
	mark bool
}

func isExisted(tree *BST, val int) bool {
	tmp := tree
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

func find(tree *BST, val int) *BST {
	tmp := tree
	for tmp != nil {
		if tmp.val == val {
			return tmp
		} else if tmp.val > val {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	return nil
}

func insert(tree *BST, val int) *BST {
	temp := &BST{val: val}

	if tree == nil {
		return temp
	}

	cur := tree
	var prev *BST
	for cur != nil {
		prev = cur
		if cur.val > temp.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	temp.parent = prev
	if prev.val > temp.val {
		prev.left = temp
	} else {
		prev.right = temp
	}
	return tree
}

func treeMinimum(tree *BST) *BST {
	t := tree
	for t.left != nil {
		t = t.left
	}
	return t
}

func treeMaximum(tree *BST) *BST {
	t := tree
	for t.right != nil {
		t = t.right
	}
	return t
}

// 寻找后继（比tree大的最小子树）
func treeSuccessor(tree *BST) *BST {
	if tree.right != nil {
		return treeMinimum(tree.right)
	}
	t := tree.parent
	for t != nil && tree == t.right {
		tree = t
		t = t.parent
	}
	return t
}

// 使用newSubTree替换subTree
func transplant(tree, subTree, newSubTree *BST) *BST {
	if subTree.parent == nil {
		tree = newSubTree
	} else if subTree == subTree.parent.left {
		subTree.parent.left = newSubTree
	} else {
		subTree.parent.right = newSubTree
	}

	if newSubTree != nil {
		newSubTree.parent = subTree.parent
	}
	return tree
}

func treeDelete(tree *BST, val int) *BST {
	target := find(tree, val)
	if target == nil {
		return tree
	}
	if target.left == nil {
		tree = transplant(tree, target, target.right)
	} else if target.right == nil {
		tree = transplant(tree, target, target.left)
	} else {
		t := treeMinimum(target.right)
		if t != target {
			tree = transplant(tree, t, t.right)
			t.right = target.right
			t.right.parent = target
		}
		transplant(tree, target, t)
		t.left = target.left
		t.left.parent = t
	}
	return tree
}
