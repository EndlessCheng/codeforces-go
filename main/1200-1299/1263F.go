package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1263F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m, v int
	Fscan(in, &n)
	maxDel := make([][]int, n+1)
	for i := range maxDel {
		maxDel[i] = make([]int, n+1)
	}
	f := func() {
		Fscan(in, &m)
		g := make([][]int, m+1)
		for w := 2; w <= m; w++ {
			Fscan(in, &v)
			g[v] = append(g[v], w)
		}
		tar := make([]int, m+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			tar[v] = i
		}
		var f func(int) (sz, l, r int)
		f = func(v int) (edgeSize, l, r int) {
			if v > 1 {
				edgeSize = 1 // 到父节点的边
			}
			l = 1e9
			if tar[v] > 0 {
				l, r = tar[v], tar[v]
			}
			for _, w := range g[v] {
				sz, ll, rr := f(w)
				edgeSize += sz
				l = min(l, ll)
				r = max(r, rr)
			}
			// 直接求区间。因为电缆没有交叉，且设备编号是连续的（见题目中的 Formally, for each tree exists a depth-first search ... 这段）
			maxDel[l][r] = max(maxDel[l][r], edgeSize)
			return
		}
		f(1)
	}
	f()
	f()

	dp := make([]int, n+1)
	for i := range dp {
		for j, dv := range dp[:i] {
			dp[i] = max(dp[i], dv+maxDel[j+1][i])
		}
	}
	Fprint(out, dp[n])
}

//func main() { CF1263F(os.Stdin, os.Stdout) }
