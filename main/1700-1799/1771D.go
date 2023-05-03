package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1771D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, v, w int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		move1 := make([][]int, n)
		for i := range move1 {
			move1[i] = make([]int, n)
		}
		for rt := range move1 {
			var build func(int, int)
			build = func(v, fa int) {
				move1[v][rt] = fa
				for _, w := range g[v] {
					if w != fa {
						build(w, v)
					}
				}
			}
			build(rt, rt)
		}

		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, n)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(int, int) int
		f = func(v, w int) (res int) {
			if v > w { // 神优化
				v, w = w, v
			}
			if v == w {
				return 1
			}
			if move1[v][w] == w {
				if s[v] == s[w] {
					return 2
				}
				return 1
			}
			dv := &dp[v][w]
			if *dv != -1 {
				return *dv
			}
			defer func() { *dv = res }()
			if s[v] == s[w] {
				return 2 + f(move1[v][w], move1[w][v])
			}
			return max(f(move1[v][w], w), f(v, move1[w][v]))
		}
		ans := 0
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				ans = max(ans, f(i, j))
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1771D(os.Stdin, os.Stdout) }
