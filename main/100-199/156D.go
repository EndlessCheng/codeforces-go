package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf156D(in io.Reader, out io.Writer) {
	var n, m, mod, cc, cnt int
	Fscan(in, &n, &m, &mod)
	g := make([][]int, n)
	for range m {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := 1
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		cnt++
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
	}
	for st, b := range vis {
		if !b {
			cc++
			cnt = 0
			dfs(st)
			ans = ans * cnt % mod
		}
	}
	if cc == 1 {
		Fprint(out, 1)
		return
	}
	for range cc - 2 {
		ans = ans * n % mod
	}
	Fprint(out, ans)
}

//func main() { cf156D(bufio.NewReader(os.Stdin), os.Stdout) }
