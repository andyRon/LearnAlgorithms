package main

// 基于环形数组实现的双向队列
type ArrayDeque struct {
	data          []interface{}
	front         int
	dequeSize     int
	dequeCapacity int
}

// 初始化双向队列
func NewArrayDeque(capacity int) *ArrayDeque {
	return &ArrayDeque{
		data:          make([]interface{}, capacity),
		front:         0,
		dequeSize:     0,
		dequeCapacity: capacity,
	}
}

// 队首入队
func (d *ArrayDeque) Push_frist(e interface{}) {
	if d.Size() == d.dequeCapacity {
		panic("队列已满")
	}
	d.front = (d.front - 1 + d.dequeCapacity) % d.dequeCapacity
	d.data[d.front] = e
	d.dequeSize++
}

// 队尾入队
func (d *ArrayDeque) Push_last(e interface{}) {
	if d.Size() == d.dequeCapacity {
		panic("队列已满")
	}
	d.data[(d.front+d.dequeSize)%d.dequeCapacity] = e
	d.dequeSize++
}

// 队首出队
func (d *ArrayDeque) Pop_first() interface{} {
	if d.IsEmpty() {
		panic("队列为空")
	}
	e := d.data[d.front]
	d.front = (d.front + 1) % d.dequeCapacity
	d.dequeSize--
	return e
}

// 队尾出队
func (d *ArrayDeque) Pop_last() interface{} {
	if d.IsEmpty() {
		panic("队列为空")
	}
	index := (d.front + d.dequeSize - 1) % d.dequeCapacity
	e := d.data[index]
	d.dequeSize--
	return e
}

// 访问队首元素
func (d *ArrayDeque) Peek_first() interface{} {
	if d.IsEmpty() {
		panic("队列为空")
	}
	return d.data[d.front]
}

// 访问队尾元素
func (d *ArrayDeque) Peek_last() interface{} {
	if d.IsEmpty() {
		panic("队列为空")
	}
	index := (d.front + d.dequeSize - 1) % d.dequeCapacity
	return d.data[index]
}

// 队列长度
func (d *ArrayDeque) Size() int {
	return d.dequeSize
}

// 判断队列是否为空
func (d *ArrayDeque) IsEmpty() bool {
	return d.dequeSize == 0
}

// 获取Slice用于打印
func (d *ArrayDeque) ToSlice() []interface{} { // TODO
	slice := make([]interface{}, d.dequeSize)
	for i := 0; i < d.dequeSize; i++ {
		slice[i] = d.data[(d.front+i)%d.dequeCapacity]
	}
	return slice
}
