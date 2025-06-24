package main

func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		// 将该最小元素与未排序区间的首个元素交换
		nums[i], nums[k] = nums[k], nums[i]
	}
}
