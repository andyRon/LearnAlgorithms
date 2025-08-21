package main

import "fmt"

// hashMapOpenAddressing 开放寻址哈希表
type hashMapOpenAddressing struct {
	buckets     []*pair
	size        int
	capacity    int
	loadThres   float64
	extendRatio int
	TOMBSTONE   *pair // 删除标记
}

func NewHashMapOpenAddressing() *hashMapOpenAddressing {
	return &hashMapOpenAddressing{
		buckets:     make([]*pair, 4),
		size:        0,
		capacity:    4,
		loadThres:   0.75,
		extendRatio: 2,
		TOMBSTONE:   &pair{-1, "-1"},
	}
}

func (h *hashMapOpenAddressing) hashFunc(key int) int {
	return key % h.capacity
}

// 计算当前负载因子
func (h *hashMapOpenAddressing) loadFactor() float64 {
	return float64(h.size) / float64(h.capacity)
}

// 搜索 key 对应的桶索引
func (h *hashMapOpenAddressing) findBucket(key int) int { // TODO
	index := h.hashFunc(key)
	firstTombstone := -1 // 记录遇到的第一个TOMBSTONE的位置
	for h.buckets[index] != nil {
		if h.buckets[index].key == key {
			if firstTombstone != -1 {
				// 若之前遇到了删除标记，则将键值对移动至该索引处
				h.buckets[firstTombstone] = h.buckets[index]
				h.buckets[index] = h.TOMBSTONE
				return firstTombstone // 返回移动后的桶索引
			}
			return index // 返回找到的索引
		}
		if firstTombstone == -1 && h.buckets[index] == h.TOMBSTONE {
			firstTombstone = index // 记录遇到的首个删除标记的位置
		}
		index = (index + 1) % h.capacity // 线性探测，越过尾部则返回头部
	}
	// 若 key 不存在，则返回添加点的索引
	if firstTombstone != -1 {
		return firstTombstone
	}
	return index
}

func (h *hashMapOpenAddressing) get(key int) string {
	index := h.findBucket(key) // 搜索 key 对应的桶索引
	if h.buckets[index] != nil && h.buckets[index] != h.TOMBSTONE {
		return h.buckets[index].val // 若找到键值对，则返回对应 val
	}
	return "" // 若键值对不存在，则返回 ""
}

// 添加操作
func (h *hashMapOpenAddressing) put(key int, val string) {
	if h.loadFactor() > h.loadThres {
		h.extend() // 当负载因子超过阈值时，执行扩容
	}
	index := h.findBucket(key) // 搜索 key 对应的桶索引
	if h.buckets[index] == nil || h.buckets[index] == h.TOMBSTONE {
		h.buckets[index] = &pair{key, val} // 若键值对不存在，则添加该键值对
		h.size++
	} else {
		h.buckets[index].val = val // 若找到键值对，则覆盖 val
	}
}

// 删除操作
func (h *hashMapOpenAddressing) remove(key int) {
	index := h.findBucket(key) // 搜索 key 对应的桶索引
	if h.buckets[index] != nil && h.buckets[index] != h.TOMBSTONE {
		h.buckets[index] = h.TOMBSTONE // 若找到键值对，则用删除标记覆盖它
		h.size--
	}
}

// 扩容哈希表
func (h *hashMapOpenAddressing) extend() {
	oldBuckets := h.buckets               // 暂存原哈希表
	h.capacity *= h.extendRatio           // 更新容量
	h.buckets = make([]*pair, h.capacity) // 初始化扩容后的新哈希表
	h.size = 0                            // 重置大小
	// 将键值对从原哈希表搬运至新哈希表
	for _, pair := range oldBuckets {
		if pair != nil && pair != h.TOMBSTONE {
			h.put(pair.key, pair.val)
		}
	}
}

// 打印哈希表
func (h *hashMapOpenAddressing) print() {
	for _, pair := range h.buckets {
		if pair == nil {
			fmt.Println("nil")
		} else if pair == h.TOMBSTONE {
			fmt.Println("TOMBSTONE")
		} else {
			fmt.Printf("%d -> %s\n", pair.key, pair.val)
		}
	}
}
