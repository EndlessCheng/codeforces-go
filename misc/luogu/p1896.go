package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func p1896(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, 1<<n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(int, int, int) int
	f = func(i, s, k int) (res int) {
		if i < 0 {
			if k > 0 {
				return
			}
			return 1
		}
		dv := &dp[i][s][k]
		if *dv >= 0 {
			return *dv
		}
		for sub, ok := s, true; ok; ok = sub != s {
			if sub<<1&sub == 0 {
				if c := bits.OnesCount(uint(sub)); c <= k {
					res += f(i-1, (1<<n-1)&^(sub<<1|sub|sub>>1), k-c)
				}
			}
			sub = (sub - 1) & s
		}
		*dv = res
		return
	}
	Fprint(out, f(n-1, 1<<n-1, k))
}

//func main() { p1896(os.Stdin, os.Stdout) }
