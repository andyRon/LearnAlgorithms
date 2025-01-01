package main

// 列表实现

// 列表类
type myList struct {
	arrCapacity int   // 列表容量
	arr         []int // 数组（存储列表元素）
	arrSize     int   // 列表长度（当前元素数量）
	extendRatio int   // 每次列表扩容的倍数
}

/* 构造函数 */
func newMyList() *myList {
	return &myList{
		arrCapacity: 10,
		arr:         make([]int, 10),
		arrSize:     0,
		extendRatio: 2,
	}
	// TODO
}
