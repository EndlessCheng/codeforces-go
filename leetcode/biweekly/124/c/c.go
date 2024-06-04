package main

// https://space.bilibili.com/206214
func maxOperations(nums []int) int {
	n := len(nums)
	res1 := helper(nums[2:n], nums[0]+nums[1])      // 删除前两个数
	res2 := helper(nums[:n-2], nums[n-2]+nums[n-1]) // 删除后两个数
	res3 := helper(nums[1:n-1], nums[0]+nums[n-1])  // 删除第一个和最后一个数
	return max(res1, res2, res3) + 1                // 加上第一次操作
}

func helper(a []int, target int) int {
	n := len(a)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if a[i]+a[i+1] == target { // 删除前两个数
				f[i][j+1] = max(f[i][j+1], f[i+2][j+1]+1)
			}
			if a[j-1]+a[j] == target { // 删除后两个数
				f[i][j+1] = max(f[i][j+1], f[i][j-1]+1)
			}
			if a[i]+a[j] == target { // 删除第一个和最后一个数
				f[i][j+1] = max(f[i][j+1], f[i+1][j]+1)
			}
		}
	}
	return f[0][n]
}

//

func maxOperations2(nums []int) int {
	n := len(nums)
	res1, done := helper2(nums[2:], nums[0]+nums[1]) // 删除前两个数
	if done {
		return n / 2
	}
	res2, done := helper2(nums[:n-2], nums[n-2]+nums[n-1]) // 删除后两个数
	if done {
		return n / 2
	}
	res3, done := helper2(nums[1:n-1], nums[0]+nums[n-1]) // 删除第一个和最后一个数
	if done {
		return n / 2
	}
	return max(res1, res2, res3) + 1 // 加上第一次操作
}

func helper2(a []int, target int) (res int, done bool) {
	n := len(a)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if done {
			return
		}
		if i >= j {
			done = true
			return
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		if a[i]+a[i+1] == target { // 删除前两个数
			res = max(res, dfs(i+2, j)+1)
		}
		if a[j-1]+a[j] == target { // 删除后两个数
			res = max(res, dfs(i, j-2)+1)
		}
		if a[i]+a[j] == target { // 删除第一个和最后一个数
			res = max(res, dfs(i+1, j-1)+1)
		}
		*p = res // 记忆化
		return
	}
	res = dfs(0, n-1)
	return
}
