package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1823F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	var n, s, t int
	Fscan(in, &n, &s, &t)
	g := make([][]int, n+1)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	f := make([]int, n+1)
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		if v == t {
			return 1
		}
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			r := dfs(w, v)
			if r > 0 {
				f[v] = r
				return r + 1
			}
		}
		return 0
	}
	dfs(s, 0)

	var dfs2 func(int, int)
	dfs2 = func(v, fa int) {
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			if w != t && f[w] == 0 {
				f[w] = f[v]
			}
			dfs2(w, v)
		}
	}
	dfs2(s, 0)

	for i := 1; i <= n; i++ {
		if i != t {
			Fprint(out, f[i]*len(g[i])%mod, " ")
		} else {
			Fprint(out, "1 ")
		}
	}
}

//func main() { cf1823F(bufio.NewReader(os.Stdin), os.Stdout) }
