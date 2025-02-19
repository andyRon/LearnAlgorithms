package main

// 列表实现

// 列表类
type myList struct {
	arrCapacity int   // 列表容量
	arr         []int // 数组（存储列表元素）
	arrSize     int   // 列表长度（当前元素数量）
	extendRatio int   // 每次列表扩容的倍数
}

// 构造函数
func newMyList() *myList {
	return &myList{
		arrCapacity: 10,
		arr:         make([]int, 10),
		arrSize:     0,
		extendRatio: 2,
	}
}

// 获取列表长度（当前元素数量）
func (l *myList) size() int {
	return l.arrSize
}

// 获取列表容量
func (l *myList) capacity() int {
	return l.arrCapacity
}

// 访问元素
func (l *myList) get(index int) int {
	if index < 0 || index >= l.arrSize {
		panic("index out of range")
	}
	return l.arr[index]
}

// 更新元素
func (l *myList) set(index int, value int) {
	if index < 0 || index >= l.arrSize {
		panic("index out of range")
	}
	l.arr[index] = value
}

// 在列表末尾添加元素
func (l *myList) add(value int) {
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	l.arr[l.arrSize] = value
	l.arrSize++
}

// 在中间插入元素
func (l *myList) insert(num, index int) {
	if index < 0 || index >= l.arrSize {
		panic("index out of range")
	}
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	// 将index以及之后的元素都向后移动一位
	for i := l.arrSize - 1; i > index; i-- {
		l.arr[i+1] = l.arr[i]
	}
	l.arr[index] = num
	l.arrSize++
}

// 删除元素
func (l *myList) remove(index int) int {
	if index < 0 || index >= l.arrSize {
		panic("index out of range")
	}
	num := l.arr[index]
	// index后元素向前移动一位
	for i := index; i < l.arrSize-1; i++ {
		l.arr[i] = l.arr[i+1]
	}
	l.arrSize--
	return num
}

// 列表扩容
func (l *myList) extendCapacity() {
	l.arr = append(l.arr, make([]int, l.arrCapacity*(l.extendRatio-1))...)
	l.arrCapacity = len(l.arr)
}

// 返回有效长度的列表
func (l *myList) toArray() []int {
	// 仅转换有效长度范围内的列表元素
	return l.arr[:l.arrSize]
}
