package main

func insertSort(nums []int) {
	l := len(nums)
	// 外循环：已排序区间为 [0, i-1]
	for i := 1; i < l; i++ {
		base := nums[i]
		j := i - 1
		// 内循环：将 base 插入到已排序区间 [0, i-1] 中的正确位置
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}
