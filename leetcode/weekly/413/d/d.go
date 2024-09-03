package main

// https://space.bilibili.com/206214
func maximumSubarrayXor(f []int, queries [][]int) []int {
	n := len(f)
	mx := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		mx[i] = make([]int, n)
		mx[i][i] = f[i]
		for j := i + 1; j < n; j++ {
			f[j] ^= f[j-1]
			mx[i][j] = max(f[j], mx[i+1][j], mx[i][j-1])
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = mx[q[0]][q[1]]
	}
	return ans
}
