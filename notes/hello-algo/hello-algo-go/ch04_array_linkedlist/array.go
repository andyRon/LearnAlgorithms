package main

import "math/rand"

func randomAccess(nums []int) (randomNum int) {
	randomIndex := rand.Intn(len(nums))
	randomNum = nums[randomIndex]
	return
}

func insert(nums []int, index int, num int) {
	for i := len(nums) - 1; i >= index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = num
}

func remove(nums []int, index int) {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

func traverse(nums []int) {
	// 通过索引遍历
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
	// 直接遍历
	for _, num := range nums {
		println(num)
	}
	for i, num := range nums {
		println(i, num)
	}
}

func find(nums []int, target int) (index int) {
	index = -1
	for i, num := range nums {
		if num == target {
			index = i
			break
		}
	}
	return
}

func extend(nums []int, enlarge int) []int {
	res := make([]int, len(nums)+enlarge)
	for i, num := range nums {
		res[i] = num
	}
	return res
}

func main() {
	nums := []int{1, 6, 10, 3, 5}
	//traverse(nums)
	println(find(nums, 10))
}
