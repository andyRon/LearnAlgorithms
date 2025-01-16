package main

// 数组表示二叉树类
type arrayBinaryTree struct {
	tree []any
}

func NewArrayBinaryTree(arr []any) *arrayBinaryTree {
	return &arrayBinaryTree{
		tree: arr,
	}
}

func (abt *arrayBinaryTree) size() int {
	return len(abt.tree)
}

// 获取索引为i节点的值
func (abt *arrayBinaryTree) val(index int) any {
	if index < 0 || index >= abt.size() {
		panic("index out of range")
	}
	return abt.tree[index]
}

// 获取索引为i节点的左子节点的索引
func (abt *arrayBinaryTree) left(i int) int {
	return 2*i + 1
}

// 获取索引为i节点的右子节点的索引
func (abt *arrayBinaryTree) right(i int) int {
	return 2*i + 2
}

// 获取索引为i节点的父节点的索引
func (abt *arrayBinaryTree) parent(i int) int {
	return (i - 1) / 2
}

// 层序遍历
func (abt *arrayBinaryTree) levelOrder() []any {
	var res []any
	for i := 0; i < abt.size(); i++ {
		if abt.val(i) != nil {
			res = append(res, abt.val(i))
		}
	}
	return res
}

// 深度优先遍历  TODO
func (abt *arrayBinaryTree) dfs(i int, order string, res *[]any) {
	if abt.val(i) == nil {
		return
	}
	if order == "pre" {
		*res = append(*res, abt.val(i))
	}
	abt.dfs(abt.left(i), order, res)
	if order == "in" {
		*res = append(*res, abt.val(i))
	}
	abt.dfs(abt.right(i), order, res)
	if order == "post" {
		*res = append(*res, abt.val(i))
	}
}

// 前序遍历
func (abt *arrayBinaryTree) preOrder() []any {
	var res []any
	abt.dfs(0, "pre", &res)
	return res
}

// 中序遍历
func (abt *arrayBinaryTree) inOrder() []any {
	var res []any
	abt.dfs(0, "in", &res)
	return res
}

// 后序遍历
func (abt *arrayBinaryTree) postOrder() []any {
	var res []any
	abt.dfs(0, "post", &res)
	return res
}
