package main

import "math"

// https://space.bilibili.com/206214
func maximumSum(nums []int, m, left, right int) int64 {
	n := len(nums)
	s := make([]int, n+1) // nums 的前缀和
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	// f[i][j] 表示在前 j 个数（下标 0 到 j-1）中选出恰好 i 个子数组，所选元素之和的最大值
	f := make([]int, n+1)
	ans := math.MinInt

	for i := 1; i <= m; i++ {
		nf := make([]int, n+1)
		for j := range nf {
			nf[j] = math.MinInt / 2
		}
		q := []int{}

		// 前 i 个子数组至少占用了 i * left 个位置
		for j := i * left; j <= n; j++ {
			// 1. 入
			k := j - left
			v := f[k] - s[k]
			for len(q) > 0 && f[q[len(q)-1]]-s[q[len(q)-1]] <= v {
				q = q[:len(q)-1]
			}
			q = append(q, k)

			// 2. 更新
			// 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
			nf[j] = max(nf[j-1], f[q[0]]-s[q[0]]+s[j])

			// 3. 出，下一轮循环队首离开窗口
			if q[0] <= j-right {
				q = q[1:]
			}
		}

		// 枚举恰好选 i 个子数组
		f = nf
		ans = max(ans, f[n])
	}

	return int64(ans)
}
