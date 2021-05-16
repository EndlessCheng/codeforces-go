package main

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7

var f = make([][]int, 1000)

func init() {
	f[0] = []int{1, 0}
	for i := 1; i < 1000; i++ {
		f[i] = make([]int, i+2)
		f[i][0] = f[i-1][0] * i % mod
		for j := 1; j <= i; j++ {
			f[i][j] = (f[i-1][j-1] + i*f[i-1][j]) % mod
		}
	}
}

func rearrangeSticks(n, k int) int {
	return f[n-1][k-1]
}
