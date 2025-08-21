package main

import (
	"fmt"
)

// HashMapChaining 链式地址哈希表
type HashMapChaining struct {
	buckets     [][]pair // 桶数组
	size        int      // 键值对数量
	capacity    int      // 哈希表容量
	loadThres   float64  // 触发扩容的负载因子阈值
	extendRatio int      // 扩容倍数
}

func NewHashMapChaining() *HashMapChaining {
	buckets := make([][]pair, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = make([]pair, 0)
	}
	return &HashMapChaining{
		buckets:     buckets,
		size:        0,
		capacity:    16,
		loadThres:   0.75,
		extendRatio: 2,
	}
}

// 哈希函数
func (m *HashMapChaining) hashFunc(key int) int {
	return key % m.capacity
}

// 负载因子
func (m *HashMapChaining) loadFactor() float64 {
	return float64(m.size) / float64(m.capacity)
}

// 查询操作
func (m *HashMapChaining) get(key int) string {
	index := m.hashFunc(key)
	for _, pair := range m.buckets[index] {
		if pair.key == key {
			return pair.val
		}
	}
	return ""
}

// 添加操作
func (m *HashMapChaining) put(key int, val string) {
	if m.loadFactor() > m.loadThres {
		m.extend()
	}
	index := m.hashFunc(key)
	// 如果有就更新
	for i := range m.buckets[index] {
		if m.buckets[index][i].key == key {
			m.buckets[index][i].val = val
			return
		}
	}
	// 如果没有就添加到尾部
	m.buckets[index] = append(m.buckets[index], pair{key: key, val: val})
	m.size++
}

// 删除操作
func (m *HashMapChaining) remove(key int) {
	index := m.hashFunc(key)
	for i := range m.buckets[index] {
		if m.buckets[index][i].key == key {
			m.buckets[index] = append(m.buckets[index][:i], m.buckets[index][i+1:]...) // TODO nice
			m.size--
			break
		}
	}
}

// 扩容哈希表
func (m *HashMapChaining) extend() {
	// 暂存原哈希表
	oldBuckets := make([][]pair, len(m.buckets))
	for i := 0; i < len(m.buckets); i++ {
		oldBuckets[i] = make([]pair, len(m.buckets[i]))
		copy(oldBuckets[i], m.buckets[i])
	}
	// 初始化扩容后的新哈希表
	m.capacity *= m.extendRatio
	m.buckets = make([][]pair, m.capacity)
	for i := 0; i < m.capacity; i++ {
		m.buckets[i] = make([]pair, 0)
	}
	m.size = 0
	// 将键值对从原哈希表搬运至新哈希表
	for _, bucket := range oldBuckets {
		for _, pair := range bucket {
			m.put(pair.key, pair.val)
		}
	}
}

// 打印哈希表
func (m *HashMapChaining) print() {
	for i := 0; i < len(m.buckets); i++ {
		fmt.Printf("%d: ", i)
		for _, pair := range m.buckets[i] {
			fmt.Printf("(%d, %s) ", pair.key, pair.val)
		}
		fmt.Println()
	}
}

//func (m *HashMapChaining) print() {
//	var builder strings.Builder
//
//	for _, bucket := range m.buckets {
//		builder.WriteString("[")
//		for _, p := range bucket {
//			builder.WriteString(strconv.Itoa(p.key) + " -> " + p.val + " ")
//		}
//		builder.WriteString("]")
//		fmt.Println(builder.String())
//		builder.Reset()
//	}
//}
