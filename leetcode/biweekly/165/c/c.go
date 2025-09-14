package main

// https://space.bilibili.com/206214
func generateSchedule(n int) [][]int {
	if n < 5 {
		return nil
	}

	ans := make([][]int, 0, n*(n-1)) // 预分配空间

	// 单独处理 d=1
	for i := 0; i < n; i += 2 {
		ans = append(ans, []int{i, (i + 1) % n})
	}
	for i := 1; i < n; i += 2 {
		ans = append(ans, []int{i, (i + 1) % n})
	}
	if n%2 == 0 { // 保证 d=1 的最后一场比赛与 d=2 的第一场比赛无冲突
		ans[len(ans)-1], ans[len(ans)-2] = ans[len(ans)-2], ans[len(ans)-1]
	}

	// 处理 d=2,3,...,n-2
	for d := 2; d < n-1; d++ {
		for i := range n {
			ans = append(ans, []int{i, (i + d) % n})
		}
	}

	// 单独处理 d=n-1
	for i := 1; i < n; i += 2 {
		ans = append(ans, []int{i, (i + n - 1) % n})
	}
	if n%2 == 0 { // 保证 a 为奇数时的最后一场比赛与 a 为偶数时的第一场比赛无冲突
		ans[len(ans)-1], ans[len(ans)-2] = ans[len(ans)-2], ans[len(ans)-1]
	}
	for i := 0; i < n; i += 2 {
		ans = append(ans, []int{i, (i + n - 1) % n})
	}

	return ans
}
