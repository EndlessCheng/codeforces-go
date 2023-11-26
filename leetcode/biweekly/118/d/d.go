package main

// https://space.bilibili.com/206214
func findMaximumLength(nums []int) (ans int) {
	n := len(nums)
	s := make([]int, n+1)
	f := make([]int, n+1)
	last := make([]int, n+1)
	q := []int{0}
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]

		// 去掉队首无用数据（计算转移直接取队首）
		for len(q) > 1 && s[q[1]]+last[q[1]] <= s[i] {
			q = q[1:]
		}

		// 计算转移
		f[i] = f[q[0]] + 1
		last[i] = s[i] - s[q[0]]

		// 去掉队尾无用数据
		for len(q) > 0 && s[q[len(q)-1]]+last[q[len(q)-1]] >= s[i]+last[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return f[n]
}
