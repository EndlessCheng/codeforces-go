package main

// https://space.bilibili.com/206214
func generateSchedule(n int) [][]int {
	if n < 5 {
		return nil
	}

	ans := make([][]int, 0, n*(n-1)) // 预分配空间

	// 处理 d=2,3,...,n-2
	for d := 2; d < n-1; d++ {
		for i := range n {
			ans = append(ans, []int{i, (i + d) % n})
		}
	}

	// 交错排列 d=1 与 d=n-1（或者说 d=-1）
	for i := range n {
		ans = append(ans, []int{i, (i + 1) % n}, []int{(i + n - 1) % n, (i + n - 2) % n})
	}

	return ans
}
