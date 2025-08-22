package ch14_dynamic_programming

// 爬楼梯，爬到楼顶方案，通过回溯来穷举所有可能性

// 回溯
func backtrack(choices []int, state, n int, res []int) {

	if state == n {
		res[0] = res[0] + 1
	}

	for _, choice := range choices {
		// TODO
	}
}
