package main

// 基础环形数组实现的队列
type ArrayQueue struct {
	data          []any // 用户存储队列元素的数组
	front         int   // 队首指针
	queueSize     int   // 队列长度
	queueCapacity int   // 队列容量
}

// 初始化队列
func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		data:          make([]any, capacity),
		front:         0,
		queueSize:     0,
		queueCapacity: capacity,
	}
}

// 入队
func (q *ArrayQueue) Push(e any) {
	if q.queueSize == q.queueCapacity {
		panic("Queue is full")
	}
	// 通过取余操作计算索引
	q.data[(q.front+q.queueSize)%q.queueCapacity] = e
	q.queueSize++
}

// 出队
func (q *ArrayQueue) Pop() any {
	if q.queueSize == 0 {
		panic("Queue is empty")
	}
	e := q.Peek()
	q.front = (q.front + 1) % q.queueCapacity // TODO
	q.queueSize--
	return e
}

// 访问队列首元素
func (q *ArrayQueue) Peek() any {
	if q.queueSize == 0 {
		panic("Queue is empty")
	}
	return q.data[q.front]
}

// 获取队列的长度
func (q *ArrayQueue) Size() int {
	return q.queueSize
}

// 判断队列是否为空
func (q *ArrayQueue) IsEmpty() bool {
	return q.queueSize == 0
}

// 获取Slice用于打印
func (q *ArrayQueue) ToSlice() []any {
	slice := make([]any, q.queueSize)
	for i := 0; i < q.queueSize; i++ {
		slice[i] = q.data[(q.front+i)%q.queueCapacity] // TODO
	}
	return slice
}
