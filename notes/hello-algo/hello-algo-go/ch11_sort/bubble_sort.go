package main

func bubbleSort(nums []int) {
	l := len(nums)
	// 外循环：未排序区间为 [0, i]
	for i := l - 1; i > 0; i-- {
		// 内循环：将未排序区间 [0, i] 中的最大元素交换至该区间的最右端
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func bubbleSortWithFlag(nums []int) {
	l := len(nums)
	for i := l - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if flag == false { // 此轮“冒泡”未交换任何元素，直接跳出
			break
		}
	}
}
