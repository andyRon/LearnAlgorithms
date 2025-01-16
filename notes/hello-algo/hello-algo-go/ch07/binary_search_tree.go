package main

type binarySearchTree struct {
	root *TreeNode
}

func newBinarySearchTree() *binarySearchTree {
	bst := &binarySearchTree{}
	// 初始化空树
	bst.root = nil
	return bst
}

// 获取根节点
func (bst *binarySearchTree) getRoot() *TreeNode {
	return bst.root
}

// TODO
