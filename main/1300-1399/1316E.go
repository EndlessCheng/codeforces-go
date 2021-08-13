package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1316E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	type pair struct {
		v int
		w []int
	}

	var n, playerNum, audienceNum int
	Fscan(in, &n, &playerNum, &audienceNum)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
	}
	for i := range a {
		a[i].w = make([]int, playerNum)
		for j := range a[i].w {
			Fscan(in, &a[i].w[j])
		}
	}
	// 排序后，满足条件就可以直接当观众
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })

	dp := make([]int64, 1<<playerNum)
	for i := range dp {
		dp[i] = -1e18
	}
	dp[0] = 0
	for i, p := range a {
		for s := 1<<playerNum - 1; s >= 0; s-- { // 倒序转移从而可以滚动数组
			if bits.OnesCount8(uint8(s)) > i-audienceNum {
				dp[s] += int64(p.v) // 直接当观众
			}
			for t, lb := s, 0; t > 0; t ^= lb {
				lb = t & -t
				dp[s] = max(dp[s], dp[s^lb]+int64(p.w[bits.TrailingZeros8(uint8(lb))]))
			}
		}
	}
	Fprint(out, dp[1<<playerNum-1])
}

//func main() { CF1316E(os.Stdin, os.Stdout) }
