package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1467E(in io.Reader, out io.Writer) {
	var n, ts, s, ans int
	Fscan(in, &n)
	a := make([]int, n)
	tot := map[int]int{}
	for i := range a {
		Fscan(in, &a[i])
		tot[a[i]]++
	}
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	d := make([]int, n+1)
	inT := make([]int, n)
	cnt := map[int]int{}
	var dfs func(int, int)
	dfs = func(v, fa int) {
		inT[v] = ts
		ts++
		color := a[v]
		c0 := cnt[color]
		cnt[color]++
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			c := cnt[color]
			dfs(w, v)
			if cnt[color] > c { // 说明 w 中有 color，那么另一侧 v（除去 w）是坏的
				// 除去子树 w 的其余点 +1（注意这不含 w）
				d[0]++
				d[inT[w]]--
				d[ts]++
			}
		}
		// 说明子树 v 是坏的
		if cnt[color]-c0 < tot[color] {
			d[inT[v]]++
			d[ts]--
		}
	}
	dfs(0, -1)

	for _, v := range d[:n] {
		s += v
		if s == 0 {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { cf1467E(bufio.NewReader(os.Stdin), os.Stdout) }
