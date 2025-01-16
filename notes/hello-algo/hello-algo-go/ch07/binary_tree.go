package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5

	// 插入
	p := NewTreeNode(0)
	n1.Left = p
	p.Left = n2
	// 删除p
	n1.Left = n2

	levelOrder(n1)
}

// 层序遍历【广度有限遍历（搜索）bfs】  借助队列来实现
func levelOrder(root *TreeNode) []any {
	// 初始化队列，加入根节点
	queue := list.New()
	queue.PushBack(root)
	//
	nums := make([]any, 0)
	for queue.Len() > 0 {
		// 队列出队
		node := queue.Remove(queue.Front()).(*TreeNode)
		// 保存节点值
		nums = append(nums, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}

// 前序、中序和后序遍历都属于深度优先遍历（depth-first traversal），也称深度优先搜索（depth-first search, DFS）
var nums []any

// 前序  根节点 -> 左子树 -> 右子树
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	nums = append(nums, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

// 中序  左子树 -> 根节点 -> 右子树
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	nums = append(nums, node.Val)
	inOrder(node.Right)
}

// 后序  左子树 -> 右子树 -> 根节点
func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	postOrder(node.Left)
	postOrder(node.Right)
	nums = append(nums, node.Val)
}
