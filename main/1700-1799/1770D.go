package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1770D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var T, n, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for _, v := range a {
			v--
			Fscan(in, &w)
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		vis := make([]bool, n)
		c := 0
		var dfs func(int, int)
		dfs = func(v, fa int) {
			vis[v] = true
			for _, w := range g[v] {
				if w != fa {
					if w == v {
						c = n
					} else if vis[w] {
						c = 2
					} else {
						dfs(w, v)
					}
				}
			}
		}
		ans := int64(1)
		for i, b := range vis {
			if !b {
				c = 0
				dfs(i, -1)
				ans = ans * int64(c) % mod
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1770D(os.Stdin, os.Stdout) }
