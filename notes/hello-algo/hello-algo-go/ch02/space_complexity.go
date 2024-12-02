package main

import "fmt"

type node struct {
	val  int
	next *node
}

func newNode(val int) *node {
	return &node{val: val}
}

func function() int {
	// 执行某些操作...
	return 0
}

func algorithm(n int) int { // 输入数据
	const a = 0      // 暂存数据（常量）
	b := 0           // 暂存数据（变量）
	newNode(0)       // 暂存数据（对象）
	c := function()  // 栈帧空间（调用函数）
	return a + b + c // 输出数据
}

func algorithm2(n int) {
	a := 0
	b := make([]int, 10000)
	var nums []int
	if n > 10 {
		nums = make([]int, n)
	}
	fmt.Println(a, b, nums)
}

// 常数阶
func spaceConstant(n int) {
	const a = 0
	b := 0
	nums := make([]int, 10000)
	node := newNode(0)

	var c int
	for i := 0; i < n; i++ {
		c = 0
	}
	for i := 0; i < n; i++ {
		function()
	}
	b += 0
	c += 0
	nums[0] = 0
	node.val = 0
}

