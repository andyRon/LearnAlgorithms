package main

import "fmt"

// 迭代

/* for 循环 */
func forLoop(n int) int {
	res := 0
	// 循环求和 1, 2, ..., n-1, n
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

/* while 循环 */
func whileLoop(n int) int {
	res := 0
	// 初始化条件变量
	i := 1
	// 循环求和 1, 2, ..., n-1, n
	for i <= n {
		res += i
		// 更新条件变量
		i++
	}
	return res
}

/* while 循环（两次更新） */
func whileLoopII(n int) int {
	res := 0
	// 初始化条件变量
	i := 1
	// 循环求和 1, 4, 10, ...
	for i <= n {
		res += i
		// 更新条件变量
		i++
		i *= 2
	}
	return res
}

func nestedForLoop(n int) string {
	res := ""
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			res += fmt.Sprintf("(%d, %d), ", i, j)
		}
	}
	return res
}

func main() {
	//fmt.Println(forLoop(5))
	//fmt.Println(whileLoop(5))
	//fmt.Println(whileLoopII(5))
	fmt.Println(nestedForLoop(5))
}
