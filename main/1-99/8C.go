package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF8C(in io.Reader, out io.Writer) {
	var n int
	a := make([]struct{ x, y int }, 25)
	Fscan(in, &a[0].x, &a[0].y, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].x, &a[i].y)
	}
	a = a[:n+1]
	dis := make([][]int, n+1)
	for i, p := range a {
		dis[i] = make([]int, n+1)
		for j, q := range a {
			dis[i][j] = (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)
		}
	}

	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = 1e9
	}
	dp[0] = 0
	from := make([]int, 1<<n)
	for s, dv := range dp {
		if dv == 1e9 { // 重要优化：由于总是选补集最小的，某些状态没访问到，也不需要从这转移
			continue
		}
		i := 1<<n - 1 ^ s
		lb := i & -i // lowbit 早晚都要选，或者说每次搬运是互相独立的
		p := bits.TrailingZeros(uint(lb)) + 1
		for j, lb2 := i, 0; j > 0; j ^= lb2 { // 可以只选 lowbit
			lb2 = j & -j
			q := bits.TrailingZeros(uint(lb2)) + 1
			if t, v := s|lb|lb2, dv+dis[0][p]+dis[0][q]+dis[p][q]; v < dp[t] {
				dp[t], from[t] = v, s
			}
		}
	}
	Fprintln(out, dp[1<<n-1])
	Fprint(out, 0)
	for i := 1<<n - 1; i > 0; i = from[i] {
		// 用 i^from[i] 求出用了哪些比特来转移
		for j, lb := i^from[i], 0; j > 0; j ^= lb {
			lb = j & -j
			Fprint(out, " ", bits.TrailingZeros(uint(lb))+1)
		}
		Fprint(out, " 0")
	}
}

//func main() { CF8C(os.Stdin, os.Stdout) }
