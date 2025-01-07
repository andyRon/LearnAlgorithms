package main

// 基于数组实现的栈
type ArrayStack struct {
	data []interface{}
}

// 初始化栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0),
	}
}

// 入栈
func (s *ArrayStack) Push(v interface{}) {
	s.data = append(s.data, v)
}

// 出栈
func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

// 访问栈顶元素
func (s *ArrayStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// 获取栈的长度
func (s *ArrayStack) Size() int {
	return len(s.data)
}

// 判断栈是否为空
func (s *ArrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

// 获取List用于打印
func (s *ArrayStack) ToSlice() []interface{} {
	return s.data
}
