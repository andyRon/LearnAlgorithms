package ch08_heap

// 建堆操作

// 大顶堆
type maxHeap struct {
	data []any
}

// 构造函数，建立空堆
func newHeap() *maxHeap {
	return &maxHeap{
		data: make([]any, 0),
	}
}

// 构造函数，根据切片建堆
func newMaxHeap(nums []any) *maxHeap {
	h := &maxHeap{data: nums}
	for i := h.parent(len(h.data) - 1); i >= 0; i-- {
		h.shiftDown(i)
	}
	return h
}

// 获取左子节点的索引
func (h *maxHeap) left(i int) int {
	return 2*i + 1
}

// 获取右子节点的索引
func (h *maxHeap) right(i int) int {
	return 2*i + 2
}

// 获取父节点的索引
func (h *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *maxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// 获取堆大小
func (h *maxHeap) size() int {
	return len(h.data)
}

// 判断堆是否为空
func (h *maxHeap) isEmpty() bool {
	return len(h.data) == 0
}

// 访问堆顶元素
func (h *maxHeap) peek() any {
	if h.isEmpty() {
		panic("heap is empty")
	}
	return h.data[0]
}

// 元素入堆
func (h *maxHeap) push(val any) {
	h.data = append(h.data, val)
	h.shiftUp(h.size() - 1)
}

// 从节点i开始，从底至顶堆化  // TODO
func (h *maxHeap) shiftUp(i int) {
	for true {
		p := h.parent(i)
		if p < 0 || h.data[i].(int) <= h.data[p].(int) {
			break
		}
		h.swap(i, p)
		// 循环向上堆化
		i = p
	}
}

func (h *maxHeap) pop() any {
	if h.isEmpty() {
		panic("heap is empty")
	}
	// 交换根节点与最右叶节点（交换首元素与尾元素）
	h.swap(0, h.size()-1)
	val := h.data[h.size()-1]
	h.data = h.data[:h.size()-1]
	h.shiftDown(0)

	return val
}

// 从节点 i 开始，从顶至底堆化
func (h *maxHeap) shiftDown(i int) {
	for true {
		// 判断节点 i, l, r 中值最大的节点，记为 max
		l, r, max := h.left(i), h.right(i), i
		if l < h.size() && h.data[l].(int) > h.data[max].(int) {
			max = l
		}
		if r < h.size() && h.data[r].(int) > h.data[max].(int) {
			max = r
		}
		if max == i {
			break
		}
		h.swap(i, max)
		// 循环向下堆化
		i = max
	}
}
