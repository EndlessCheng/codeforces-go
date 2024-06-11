package main

// https://space.bilibili.com/206214
func maxScore(nums []int, x int) int64 {
	// 翻译 py 的时候改一下变量名 x
	n := len(nums)
	memo := make([][2]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1e18
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == n {
			return
		}
		p := &memo[i][j]
		if *p != -1e18 {
			return *p
		}
		defer func() { *p = res }()
		res = dfs(i+1, nums[i]%2) + nums[i]
		if nums[i]%2 != j {
			res -= x
		}
		res = max(res, dfs(i+1, j))
		return
	}
	return int64(nums[0] + dfs(1, nums[0]%2))
}
