package main

// github.com/EndlessCheng/codeforces-go
func countMaxOrSubsets(nums []int) (ans int) {
	totalOr := 0
	for _, x := range nums {
		totalOr |= x
	}
	n := len(nums)
	var dfs func(int, int)
	dfs = func(i, subsetOr int) {
		if subsetOr == totalOr {
			ans += 1 << (n - i)
			return
		}
		if i == n {
			return
		}
		dfs(i+1, subsetOr)         // 不选 nums[i]
		dfs(i+1, subsetOr|nums[i]) // 选 nums[i]
	}
	dfs(0, 0)
	return
}
