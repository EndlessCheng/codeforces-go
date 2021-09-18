package main

/* 动态规划

定义 $f[i]$ 表示行驶到 i 时的最大盈利。考虑状态转移，一方面，我们可以不接终点为 $i$ 的乘客，这样有 $f[i]=f[i-1]$；另一方面，我们可以接所有终点为 $i$ 的乘客，这样有 $f[i] = \max f[start]+end-start+tip$，二者取最大值。

最终答案为 $f[n]$。

*/

// github.com/EndlessCheng/codeforces-go
func maxTaxiEarnings(n int, rides [][]int) int64 {
	f := make([]int, n+1)
	groups := make([][][2]int, n+1)
	for _, r := range rides {
		start, end, tip := r[0], r[1], r[2]
		groups[end] = append(groups[end], [2]int{start, tip})
	}
	for end := 1; end <= n; end++ {
		f[end] = f[end-1]
		for _, r := range groups[end] {
			start, tip := r[0], r[1]
			f[end] = max(f[end], f[start]+end-start+tip)
		}
	}
	return int64(f[n])
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
