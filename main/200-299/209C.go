package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf209C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, ans, cc, even int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	deg := make([]int, n)
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		deg[v]++
		deg[w]++
	}

	vis := make([]bool, n)
	var dfs func(int) bool
	dfs = func(v int) bool {
		vis[v] = true
		odd := deg[v]%2 > 0
		for _, w := range g[v] {
			if !vis[w] && dfs(w) {
				odd = true
			}
		}
		return odd
	}
	for i, d := range deg {
		if i > 0 && d == 0 { // 必须从 0 出发，所以 0 一定要 check
			continue
		}
		if !vis[i] {
			cc++
			if !dfs(i) {
				even++
			}
		}
		ans += d % 2
	}
	ans /= 2
	if cc > 1 {
		ans += even
	}
	Fprint(out, ans)
}

//func main() { cf209C(os.Stdin, os.Stdout) }
