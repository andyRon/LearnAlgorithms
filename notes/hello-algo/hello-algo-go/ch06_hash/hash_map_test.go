package main

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	/* 初始化哈希表 */
	hmap := make(map[int]string)

	/* 添加操作 */
	// 在哈希表中添加键值对 (key, value)
	hmap[12836] = "小哈"
	hmap[15937] = "小啰"
	hmap[16750] = "小算"
	hmap[13276] = "小法"
	hmap[10583] = "小鸭"

	/* 查询操作 */
	// 向哈希表中输入键 key ，得到值 value
	name := hmap[15937]
	println(name)

	/* 删除操作 */
	// 在哈希表中删除键值对 (key, value)
	delete(hmap, 10583)

	/* 遍历哈希表 */
	// 遍历键值对 key->value
	for key, value := range hmap {
		fmt.Println(key, "->", value)
	}
	// 单独遍历键 key
	for key := range hmap {
		fmt.Println(key)
	}
	// 单独遍历值 value
	for _, value := range hmap {
		fmt.Println(value)
	}
}
