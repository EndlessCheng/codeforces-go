package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF623B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var n int
	var rm, ch int64
	Fscan(in, &n, &rm, &ch)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := int64(1e18)
	vis := map[int]bool{}
	do := func(p int) {
		if vis[p] {
			return
		}
		vis[p] = true
		dp := [3]int64{}
		// 对于删除操作，删到只剩一个肯定比全删更优，所以不会让最小值取到删全部的情况上
		for _, v := range a {
			cost := int64(0)
			if v%p == 1 || v%p == p-1 {
				cost = ch
			} else if v%p > 0 {
				cost = 1e18
			}
			dp[2] = min(min(dp[1], dp[2])+cost, 1e18)
			dp[1] = min(min(dp[0], dp[1])+rm, 1e18)
			dp[0] = min(dp[0]+cost, 1e18)
		}
		ans = min(ans, min(min(dp[0], dp[1]), dp[2]))
	}
	f := func(x int) {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				for x /= i; x%i == 0; x /= i {
				}
				do(i)
			}
		}
		if x > 1 {
			do(x)
		}
	}
	f(a[0] - 1)
	f(a[0])
	f(a[0] + 1)
	f(a[n-1] - 1)
	f(a[n-1])
	f(a[n-1] + 1)
	Fprint(out, ans)
}

//func main() { CF623B(os.Stdin, os.Stdout) }
