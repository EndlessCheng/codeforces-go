package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int = 1e9

	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	if n == 1 {
		Fprint(out, 0)
		return
	}
	ws := make([][]int, n)
	for i := range ws {
		ws[i] = make([]int, n)
		for j := range ws[i] {
			ws[i][j] = inf
		}
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		wt = min(wt, ws[v][w])
		ws[v][w] = wt
		ws[w][v] = wt
	}

	m = 1 << n
	expand := make([]int, m) // 从一个节点集合往下挖一层后的所有节点集合（旧+新）
	sDis := make([][]int, m) // 从一个节点集合出发到某一节点的最短路
	for s := range expand {
		expand[s] = s
		sDis[s] = make([]int, n)
		for j := range sDis[s] {
			sDis[s][j] = inf
		}
		for _s := uint(s); _s > 0; _s &= _s - 1 {
			v := bits.TrailingZeros(_s)
			sDis[s][v] = 0
			for w, wt := range ws[v] {
				if wt < inf {
					expand[s] |= 1 << w
					sDis[s][w] = min(sDis[s][w], wt)
				}
			}
		}
	}

	type pair struct{ set, disSum int }
	from := make([][]pair, m) // 当前节点集合可以从哪些节点集合转移过来（下挖一层），新修的路径长度和为多少
	for s := range from {
		sub := s
		for ok := true; ok; ok = sub != s {
			if s&expand[sub] == s {
				ds := 0
				for newS := uint(s ^ sub); newS > 0; newS &= newS - 1 {
					ds += sDis[sub][bits.TrailingZeros(newS)]
				}
				from[s] = append(from[s], pair{sub, ds})
			}
			sub = (sub - 1) & s
		}
	}

	ans := inf
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	for i := 0; i < n; i++ {
		dp[0][1<<i] = 0
	}
	for i := 1; i <= n; i++ {
		for j, ps := range from {
			for _, p := range ps {
				dp[i][j] = min(dp[i][j], dp[i-1][p.set]+p.disSum*i)
			}
		}
		ans = min(ans, dp[i][m-1])
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
