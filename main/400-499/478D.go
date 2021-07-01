package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF478D(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var r, g, ans int
	Fscan(in, &r, &g)
	h := 1
	for ; (h+1)*(h+2)/2 <= r+g; h++ {
	}
	dp := make([]int, r+1)
	dp[0] = 1
	for i := 1; i <= h; i++ {
		for j := r; j >= i; j-- {
			dp[j] = (dp[j] + dp[j-i]) % mod
		}
	}
	// 此时 dp[i] 表示堆了 h 层且用了 i 个红色方块的方案数
	// 为保证用的绿色方块数不超过 g，i 不能小于 max(0, h*(h+1)/2-g)
	for _, v := range dp[max(0, h*(h+1)/2-g) : r+1] {
		ans = (ans + v) % mod
	}
	Fprint(out, ans)
}

//func main() { CF478D(os.Stdin, os.Stdout) }
