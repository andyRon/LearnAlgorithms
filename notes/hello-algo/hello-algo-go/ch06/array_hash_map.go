package main

import "fmt"

// 基于数组实现的哈希表

// 键值对
type pair struct {
	key int
	val string
}

type arrayHashMap struct {
	buckets []*pair
}

// 初始化哈希表
func newArrayHashMap() *arrayHashMap {
	return &arrayHashMap{
		buckets: make([]*pair, 100),
	}
}

// 哈希函数
func (a *arrayHashMap) hashFunc(key int) int {
	return key % len(a.buckets)
}

// 查询操作
func (a *arrayHashMap) get(key int) string {
	index := a.hashFunc(key)
	if a.buckets[index] == nil {
		return "Not Found"
	}
	return a.buckets[index].val
}

// 添加操作
func (a *arrayHashMap) put(key int, val string) {
	index := a.hashFunc(key)
	a.buckets[index] = &pair{
		key: key,
		val: val,
	}
}

// 删除操作
func (a *arrayHashMap) remove(key int) {
	index := a.hashFunc(key)
	a.buckets[index] = nil
}

// 获取所有键对
func (a *arrayHashMap) pairSet() []*pair {
	var pairs []*pair
	for _, pair := range a.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

// 获取所有键
func (a *arrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

// 获取所有值
func (a *arrayHashMap) valSet() []string {
	var vals []string
	for _, pair := range a.buckets {
		if pair != nil {
			vals = append(vals, pair.val)
		}
	}
	return vals
}

// 打印哈希表
func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.val)
		}
	}
}
