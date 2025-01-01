package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestList(t *testing.T) {
	/* 初始化列表 */
	// 无初始值
	//nums1 := []int{}
	// 有初始值
	nums := []int{1, 3, 2, 5, 4}

	/* 访问元素 */
	//num := nums[1] // 访问索引 1 处的元素

	/* 更新元素 */
	nums[1] = 0 // 将索引 1 处的元素更新为 0

	/* 清空列表 */
	nums = nil

	/* 在尾部添加元素 */
	nums = append(nums, 1)
	nums = append(nums, 3)
	nums = append(nums, 2)
	nums = append(nums, 5)
	nums = append(nums, 4)

	/* 在中间插入元素 */
	nums = append(nums[:3], append([]int{6}, nums[3:]...)...) // 在索引 3 处插入数字 6   // TODO
	fmt.Println(nums)

	/* 删除元素 */
	nums = append(nums[:3], nums[4:]...) // 删除索引 3 处的元素
	fmt.Println(nums)

	/* 通过索引遍历列表 */
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	/* 直接遍历列表元素 */
	count = 0
	for _, num := range nums {
		count += num
	}

	/* 拼接两个列表 */
	nums1 := []int{6, 8, 7, 10, 9}
	nums = append(nums, nums1...) // 将列表 nums1 拼接到 nums 之后
	fmt.Println(nums)

	/* 排序列表 */
	sort.Ints(nums) // 排序后，列表元素从小到大排列
}
