package ch10_searching

// 二分查找的插入点

// 二分查找插入点（无重复元素）
func binarySearchInsertionSimple(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	// 未找到 target ，返回插入点left
	return left
}

// 二分查找插入点（存在重复元素）
func binarySearchInsertion(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid - 1 // TODO
		}
	}
	return left
}
