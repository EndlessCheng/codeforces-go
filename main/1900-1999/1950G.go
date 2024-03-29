package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1950G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		a := make([]struct{ x, y int }, n)
		idx := map[string]int{}
		for i := range a {
			Fscan(in, &s, &t)
			if idx[s] == 0 {
				idx[s] = len(idx) + 1
			}
			a[i].x = idx[s] - 1
			if idx[t] == 0 {
				idx[t] = len(idx) + 1
			}
			a[i].y = idx[t] - 1
			for j, p := range a[:i] {
				if p.x == a[i].x || p.y == a[i].y {
					g[i] = append(g[i], j)
					g[j] = append(g[j], i)
				}
			}
		}

		dp := make([][16]int, 1<<n)
		for i := range dp {
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(int, int) int
		f = func(mask, v int) (res int) {
			if mask == 1<<n-1 {
				return 1
			}
			p := &dp[mask][v]
			if *p != -1 {
				return *p
			}
			for _, w := range g[v] {
				if mask>>w&1 == 0 {
					res = max(res, f(mask|1<<w, w))
				}
			}
			res++
			*p = res
			return
		}
		ans := 0
		for i := range g {
			ans = max(ans, f(1<<i, i))
		}
		Fprintln(out, n-ans)
	}
}

//func main() { cf1950G(os.Stdin, os.Stdout) }
