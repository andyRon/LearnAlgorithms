package main

import "container/list"

// 基于链表实现的栈
type linkedListStack struct {
	// 使用内置包list来实现栈
	data *list.List
}

// 初始化栈
func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(),
	}
}

// 入栈
func (s *linkedListStack) push(v interface{}) {
	s.data.PushBack(v)
}

// 出栈
func (s *linkedListStack) pop() interface{} {
	if s.data.Len() == 0 {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

// 访问栈顶元素
func (s *linkedListStack) peek() interface{} {
	if s.data.Len() == 0 {
		return nil
	}
	return s.data.Back().Value
}

// 获取栈的长度
func (s *linkedListStack) size() int {
	return s.data.Len()
}

// 判断栈是否为空
func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

// 获取List用于打印
func (s *linkedListStack) toList() *list.List {
	return s.data
}
