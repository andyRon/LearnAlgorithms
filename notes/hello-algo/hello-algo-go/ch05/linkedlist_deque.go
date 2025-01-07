package main

import "container/list"

// 基于双向链表实现的双向队列
type Deque struct {
	data *list.List
}

// 初始化双向队列
func NewDeque() *Deque {
	return &Deque{
		data: list.New(),
	}
}

// 队首入队
func (d *Deque) Push_first(value interface{}) {
	d.data.PushFront(value)
}

// 队尾入队
func (d *Deque) Push_last(value interface{}) {
	d.data.PushBack(value)
}

// 队首出队
func (d *Deque) Pop_first() interface{} {
	if d.IsEmpty() {
		return nil
	}
	return d.data.Remove(d.data.Front())
}

// 队尾出队
func (d *Deque) Pop_last() interface{} {
	if d.IsEmpty() {
		return nil
	}
	return d.data.Remove(d.data.Back())
}

// 访问队首元素
func (d *Deque) Peek_first() interface{} {
	if d.IsEmpty() {
		return nil
	}
	return d.data.Front().Value
}

// 访问队尾元素
func (d *Deque) Peek_last() interface{} {
	if d.IsEmpty() {
		return nil
	}
	return d.data.Back().Value
}

// 队列长度
func (d *Deque) Size() int {
	return d.data.Len()
}

// 判断队列是否为空
func (d *Deque) IsEmpty() bool {
	return d.data.Len() == 0
}

// 获取List用于打印
func (d *Deque) ToList() *list.List {
	return d.data
}
