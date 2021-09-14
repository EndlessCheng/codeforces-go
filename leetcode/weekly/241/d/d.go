package main

// 我的题解：https://leetcode-cn.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/solution/zhuan-huan-cheng-di-yi-lei-si-te-lin-shu-2y1k/

// github.com/EndlessCheng/codeforces-go
var f [1001][1001]int

func init() {
	f[0][0] = 1
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= i; j++ {
			f[i][j] = (f[i-1][j-1] + (i-1)*f[i-1][j]) % (1e9 + 7)
		}
	}
}

func rearrangeSticks(n, k int) int {
	return f[n][k]
}
