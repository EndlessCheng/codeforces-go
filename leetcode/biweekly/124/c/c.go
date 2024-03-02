package main

// https://space.bilibili.com/206214
func maxOperations(nums []int) int {
	n := len(nums)
	res1 := helper(nums[2:], nums[0]+nums[1])       // 最前面两个
	res2 := helper(nums[:n-2], nums[n-1]+nums[n-2]) // 最后两个
	res3 := helper(nums[1:n-1], nums[0]+nums[n-1])  // 第一个和最后一个
	return max(res1, res2, res3) + 1                // 加上第一次操作
}

func helper(a []int, target int) int {
	n := len(a)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i >= j {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if a[i]+a[i+1] == target { // 最前面两个
			res = max(res, dfs(i+2, j)+1)
		}
		if a[j-1]+a[j] == target { // 最后两个
			res = max(res, dfs(i, j-2)+1)
		}
		if a[i]+a[j] == target { // 第一个和最后一个
			res = max(res, dfs(i+1, j-1)+1)
		}
		*p = res
		return
	}
	return dfs(0, n-1)
}
