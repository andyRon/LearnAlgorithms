package ch08

import (
	"container/heap"
	"fmt"
	"testing"
)

type intHeap []any

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	return (*h)[i].(int) > (*h)[j].(int)
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Top() any {
	return (*h)[0]
}

func TestHeap(t *testing.T) {
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	// 元素入队
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 5)

	// 获取堆顶元素
	top := maxHeap.Top()
	fmt.Printf("堆顶元素为：%d\n", top)

	// 堆顶元素出堆
	heap.Pop(maxHeap)
	heap.Pop(maxHeap)
	heap.Pop(maxHeap)
	heap.Pop(maxHeap)
	heap.Pop(maxHeap)

	// 获取堆大小
	size := len(*maxHeap)
	fmt.Printf("堆元素数量为 %d\n", size)

	// 判断堆是否为空
	isEmpty := len(*maxHeap) == 0
	fmt.Printf("堆是否为空 %t\n", isEmpty)
}