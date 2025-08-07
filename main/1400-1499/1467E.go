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
		x := a[v]
		c0 := cnt[x]
		cnt[x]++
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			c := cnt[x]
			dfs(w, v)
			if cnt[x] > c {
				d[0]++
				d[inT[w]]--
				d[ts]++
			}
		}
		if cnt[x]-c0 < tot[x] {
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
