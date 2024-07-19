package main

import "container/list"

// 递归
func recur(n int) int {
	// 终止条件
	if n == 1 {
		return 1
	}
	// 递：递归调用
	res := recur(n - 1)
	// 归：返回结果
	return n + res
}

// 尾递归
func tailRecur(n int, res int) int {
	// 终止条件
	if n == 0 {
		return res
	}
	// 尾递归调用
	return tailRecur(n-1, res+n)
}

// 斐波那契数列：递归
func fib(n int) int {
	// 终止条件
	if n == 1 || n == 2 {
		return n - 1
	}
	// 递归调用 f(n) = f(n-1) + f(n-2)
	res := fib(n-1) + fib(n-2)
	// 返回结果 f(n)
	return res
}

// 使用迭代模拟递归
func forLoopRecur(n int) int {
	// 使用一个显式的栈来模拟系统调用栈
	stack := list.New()
	res := 0
	// 递：递归调用
	for i := n; i > 0; i-- {
		// 通过入栈操作模拟"递"
		stack.PushBack(i)
	}
	// 归：返回结果
	for stack.Len() != 0 {
		// 通过出栈操作模拟"归"
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}
	return res
}
