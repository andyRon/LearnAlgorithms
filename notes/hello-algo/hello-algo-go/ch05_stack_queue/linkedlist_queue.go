package main

import "container/list"

// 基于链表实现的队列
type LinkedListQueue struct {
	// 使用内置包list来实现队列
	data *list.List
}

// 初始化队列
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		data: list.New(),
	}
}

// 入队
func (q *LinkedListQueue) Push(v interface{}) {
	q.data.PushBack(v)
}

// 出队
func (q *LinkedListQueue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.data.Remove(q.data.Front())
}

// 访问队列首元素
func (q *LinkedListQueue) Peek() interface{} {
	if q.data.Len() == 0 {
		return nil
	}
	return q.data.Front().Value
}

// 获取队列的长度
func (q *LinkedListQueue) Size() int {
	return q.data.Len()
}

// 判断队列是否为空
func (q *LinkedListQueue) IsEmpty() bool {
	return q.data.Len() == 0
}

// 获取List用于打印
func (q *LinkedListQueue) ToList() *list.List {
	return q.data
}
