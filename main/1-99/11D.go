package main

import (
	. "fmt"
	"io"
)

// https://blog.csdn.net/fangzhenpeng/article/details/49078233

// github.com/EndlessCheng/codeforces-go
func CF11D(in io.Reader, out io.Writer) {
	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := int64(0)
	dp := make([][]int64, 1<<n)
	for i := range dp {
		dp[i] = make([]int64, n)
	}
	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1
	}
	for s := range dp {
		for v, dv := range dp[s] {
			if dv == 0 {
				continue
			}
			for _, w := range g[v] {
				if 1<<w < s&-s {
					continue
				}
				if 1<<w&s == 0 {
					dp[s|1<<w][w] += dv
				} else if 1<<w == s&-s {
					ans += dv
				}
			}
		}
	}
	Fprint(out, (ans-int64(m))/2)
}

//func main() { CF11D(os.Stdin, os.Stdout) }
