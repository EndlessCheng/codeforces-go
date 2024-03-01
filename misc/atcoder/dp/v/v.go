package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w int
	type data int
	var mod data
	Fscan(in, &n, &mod)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	const root = 0
	const unitData data = 1
	moveEdge := func(d data, v, w int) data { return d + 1 }
	merge := func(p, q data) data { return p * q % mod }

	subRes := make([]data, len(g))
	var dfs func(int, int)
	dfs = func(v, fa int) {
		res := unitData
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v)
				r := moveEdge(subRes[w], v, w)
				res = merge(res, r)
			}
		}
		subRes[v] = res
	}
	dfs(root, -1)

	ans := make([]data, len(g))
	var reroot func(int, int, data)
	reroot = func(v, fa int, movedFa data) {
		if fa < 0 {
			ans[v] = subRes[v]
		} else {
			ans[v] = merge(movedFa, subRes[v])
		}

		m := len(g[v])
		suf := make([]data, m+1)
		suf[m] = unitData
		for i := m - 1; i >= 0; i-- {
			w := g[v][i]
			if w != fa {
				r := moveEdge(subRes[w], v, w)
				suf[i] = merge(suf[i+1], r)
			} else {
				suf[i] = suf[i+1]
			}
		}

		pre := unitData
		for i, w := range g[v] {
			if w != fa {
				mergeRes := merge(movedFa, merge(pre, suf[i+1]))
				reroot(w, v, moveEdge(mergeRes, w, v))
				r := moveEdge(subRes[w], v, w)
				pre = merge(pre, r)
			}
		}
	}
	reroot(root, -1, unitData)
	
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }
