package main

import "testing"

func TestStack(t *testing.T) {
	/* 初始化栈 */
	// 在 Go 中，推荐将 Slice 当作栈来使用
	var stack []int

	/* 元素入栈 */
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)
	stack = append(stack, 5)
	stack = append(stack, 4)

	/* 访问栈顶元素 */
	peek := stack[len(stack)-1]

	/* 元素出栈 */
	pop := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	/* 获取栈的长度 */
	size := len(stack)

	/* 判断是否为空 */
	isEmpty := len(stack) == 0
}
