package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	dfn := make([]int, n)
	ts := 0
	pa := make([][18]int, n)
	dep := make([]int, n)
	var build func(int, int)
	build = func(v, p int) {
		dfn[v] = ts
		ts++
		pa[v][0] = p
		for _, w := range g[v] {
			if w != p {
				dep[w] = dep[v] + 1
				build(w, v)
			}
		}
	}
	build(0, -1)
	mx := bits.Len(uint(n))
	for i := range mx - 1 {
		for v := range pa {
			p := pa[v][i]
			if p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v, d int) int {
		for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros32(k)]
		}
		return v
	}
	getLCA := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			pv, pw := pa[v][i], pa[w][i]
			if pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}

	nodesMap := map[int][]int{}
	for i, x := range a {
		nodesMap[x] = append(nodesMap[x], i)
	}

	vt := make([][]int, n)
	isNode := make([]int, n)
	for i := range isNode {
		isNode[i] = -1
	}
	addEdge := func(v, w int) {
		vt[v] = append(vt[v], w)
	}
	const root = 0
	st := []int{root}

	for t, nodes := range nodesMap {
		slices.SortFunc(nodes, func(a, b int) int { return dfn[a] - dfn[b] })
		vt[root] = vt[root][:0]
		st = st[:1]
		for _, v := range nodes {
			isNode[v] = t
			if v == root {
				continue
			}
			vt[v] = vt[v][:0]
			lca := getLCA(st[len(st)-1], v)
			for len(st) > 1 && dfn[lca] <= dfn[st[len(st)-2]] {
				addEdge(st[len(st)-2], st[len(st)-1])
				st = st[:len(st)-1]
			}
			if lca != st[len(st)-1] {
				vt[lca] = vt[lca][:0]
				addEdge(lca, st[len(st)-1])
				st[len(st)-1] = lca
			}
			st = append(st, v)
		}
		for i := 1; i < len(st); i++ {
			addEdge(st[i-1], st[i])
		}

		var dfs func(int) int
		dfs = func(v int) (size int) {
			if isNode[v] == t {
				size = 1
			}
			for _, w := range vt[v] {
				sz := dfs(w)
				wt := dep[w] - dep[v]
				ans += wt * sz * (len(nodes) - sz)
				size += sz
			}
			return
		}

		rt := root
		if isNode[rt] != t && len(vt[rt]) == 1 {
			rt = vt[rt][0]
		}
		dfs(rt)
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
