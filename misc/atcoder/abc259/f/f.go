package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, wt int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	var dfs func(int, int) (int, int)
	dfs = func(v, fa int) (int, int) {
		sumF := 0
		diff := sort.IntSlice{0}
		for _, e := range g[v] {
			if e.to != fa {
				full, notFull := dfs(e.to, v)
				sumF += full
				diff = append(diff, max(e.wt+notFull-full, 0))
			}
		}
		sort.Sort(sort.Reverse(diff))
		d := a[v]
		if d == 0 {
			return sumF, -1e18 // 禁止父节点与 v 连边
		}
		extra := 0
		for _, d := range diff[:d-1] {
			extra += d
		}
		return sumF + extra + diff[d-1], sumF + extra
	}
	full, _ := dfs(0, -1)
	Fprint(out, full)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
