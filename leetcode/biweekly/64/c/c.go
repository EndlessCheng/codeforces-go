package main

// github.com/EndlessCheng/codeforces-go
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	sum := make([]int, n+1) // sum[i] 表示 s[:i] 中盘子的个数
	left := make([]int, n)  // left[i] 表示 i 左侧最近蜡烛位置
	p := -1
	for i, b := range s {
		sum[i+1] = sum[i]
		if b == '|' {
			p = i
		} else {
			sum[i+1]++
		}
		left[i] = p
	}

	right := make([]int, n) // right[i] 表示 i 右侧最近蜡烛位置
	for i, p := n-1, n; i >= 0; i-- {
		if s[i] == '|' {
			p = i
		}
		right[i] = p
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := right[q[0]], left[q[1]]
		if l < r {
			ans[i] = sum[r] - sum[l]
		}
	}
	return ans
}
