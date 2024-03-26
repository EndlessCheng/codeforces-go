package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf19E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, ts, odd int
	Fscan(in, &n, &m)
	type edge struct{ to, i int }
	g := make([][]edge, n)
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], edge{w, i})
		g[w] = append(g[w], edge{v, i})
	}

	dfn := make([]int, n)
	dep := make([]int, n)
	diffV := make([]int, n)
	diffE := make([]int, m)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		ts++
		dfn[v] = ts
		for _, e := range g[v] {
			w := e.to
			if dfn[w] == 0 {
				dep[w] = dep[v] ^ 1
				dfs(w, v)
				diffE[e.i] = diffV[w]
				diffV[v] += diffV[w]
			} else if w != fa && dfn[w] < dfn[v] { // 返祖边
				t := 1 - (dep[w]^dep[v])*2
				if t > 0 { // 奇环
					odd++
					diffE[e.i] = 1
				}
				diffV[v] += t
				diffV[w] -= t
			}
		}
	}
	for i, d := range dfn {
		if d == 0 {
			dfs(i, -1)
		}
	}

	ans := []any{}
	for i, c := range diffE {
		if odd == 0 || c == odd {
			ans = append(ans, i+1)
		}
	}
	Fprintln(out, len(ans))
	Fprintln(out, ans...)
}

//func main() { cf19E(os.Stdin, os.Stdout) }
