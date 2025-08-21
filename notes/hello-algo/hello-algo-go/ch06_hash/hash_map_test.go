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

func TestHashMapChaining(t *testing.T) { // TODO
	hmap := NewHashMapChaining()
	hmap.put(200, "A")
	hmap.put(937, "B")
	hmap.put(750, "C")
	//hmap.put(276, "D")
	//hmap.put(583, "E")
	//hmap.put(500, "F")
	//hmap.put(300, "J")
	//hmap.put(637, "G")
	//hmap.put(437, "K")
	//hmap.put(283, "L")
	//hmap.put(237, "M")
	//hmap.put(850, "H")
	//hmap.put(683, "I")

	hmap.print()
	//for i, pair := range hmap.buckets {
	//	if pair != nil {
	//		for j, p := range pair {
	//			println(i, j, p.key, p.val)
	//		}
	//	}
	//}
}

func TestHashMapOpenAddressing(t *testing.T) {
	hmap := NewHashMapOpenAddressing()
	hmap.put(200, "A")
	hmap.put(937, "B")
	hmap.put(750, "C")
	hmap.put(276, "D")
	hmap.put(583, "E")
	hmap.put(500, "F")
	hmap.put(300, "J")
	hmap.put(637, "G")
	hmap.put(437, "K")
	hmap.put(283, "L")
	hmap.put(237, "M")
	hmap.put(850, "H")
	hmap.put(683, "I")

	//hmap.print()
	//println("------")
	//hmap.remove(200)
	//hmap.print()

	for i, pair := range hmap.buckets {
		if pair != nil && pair != hmap.TOMBSTONE {
			println(i, pair.key, pair.val)
		}
	}

}
