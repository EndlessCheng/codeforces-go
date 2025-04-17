package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (s int) {
	for ; i > 0; i &= i - 1 {
		s += f[i]
	}
	return
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, ts int
	Fscan(in, &n)
	type edge struct{ v, w, wt int }
	es := make([]edge, n-1)
	g := make([][]int, n)
	for i := range es {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		es[i] = edge{v, w, wt}
	}

	inT := make([]int, n+1)
	outT := make([]int, n+1)
	const mx = 18
	pa := make([][mx]int, n)
	dep := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, p int) {
		ts++
		inT[v] = ts
		pa[v][0] = p
		for _, w := range g[v] {
			if w == p {
				continue
			}
			dep[w] = dep[v] + 1
			dfs(w, v)
		}
		outT[v] = ts
	}
	dfs(0, -1)
	for i := 0; i < mx-1; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v, d int) int {
		for k := uint(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros(k)]
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
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}

	weight := make([]int, n+1)
	diff := make(fenwick, n+1)
	update := func(x, y, w int) {
		if inT[x] > inT[y] {
			y = x
		}
		d := w - weight[y]
		weight[y] = w
		diff.update(inT[y], d)
		diff.update(outT[y]+1, -d)
	}
	for _, e := range es {
		update(e.v, e.w, e.wt)
	}

	for Fscan(in, &q); q > 0; q-- {
		var v, w int
		Fscan(in, &op, &v, &w)
		v--
		if op == 1 {
			update(es[v].v, es[v].w, w)
		} else {
			w--
			Fprintln(out, diff.pre(inT[v])+diff.pre(inT[w])-diff.pre(inT[getLCA(v, w)])*2)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
