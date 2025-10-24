package main

import (
	"math"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func largestNumber(cost []int, target int) string {
	f := make([]int, target+1)
	for i := range f {
		f[i] = math.MinInt
	}
	f[0] = 0

	// 计算最长长度
	for _, c := range cost {
		for j := c; j <= target; j++ {
			f[j] = max(f[j], f[j-c]+1)
		}
	}
	if f[target] < 0 { // 无解
		return "0"
	}

	ans := make([]byte, 0, f[target]) // 预分配空间
	j := target
	for i, c := range slices.Backward(cost) { // 从大到小填数字
		// f[j-c]+1 == f[j] 说明上面 f[j] 取的 max 来自 f[j-c]，说明可以填 i+1
		for j >= c && f[j-c]+1 == f[j] {
			ans = append(ans, '1'+byte(i))
			j -= c
		}
	}
	return string(ans)
}
