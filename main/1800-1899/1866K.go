package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"math/big"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type vec66 struct{ x, y int }

func (a vec66) sub(b vec66) vec66 { return vec66{a.x - b.x, a.y - b.y} }
func (a vec66) dot(b vec66) int   { return a.x*b.x + a.y*b.y }
func (a vec66) detCmp(b vec66) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

func cf1866K(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, Q, ans int
	Fscan(in, &n)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for range n - 1 {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	nodes := make([]struct{ fi, se, fiW int }, n)
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		p := &nodes[v]
		for _, e := range g[v] {
			w := e.to
			if w == fa {
				continue
			}
			d := dfs(w, v) + e.wt
			ans = max(ans, p.fi+d)
			if d > p.fi {
				p.se = p.fi
				p.fi = d
				p.fiW = w
			} else if d > p.se {
				p.se = d
			}
		}
		return p.fi
	}
	dfs(0, -1)

	hulls := make([][2][]vec66, n)
	var reroot func(int, int, vec66)
	reroot = func(v, fa int, up vec66) {
		a := make([]vec66, len(g[v]))
		for i, e := range g[v] {
			w := e.to
			if w == fa {
				a[i] = up
			} else {
				a[i] = vec66{e.wt, nodes[w].fi}
			}
		}

		f := func(a, b vec66) int { return cmp.Or(a.x-b.x, a.y-b.y) }
		slices.SortFunc(a, f)
		q := a[:0]
		b := []vec66{}
		for _, v := range a {
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).detCmp(v.sub(q[len(q)-1])) >= 0 {
				b = append(b, q[len(q)-1])
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
		hulls[v][0] = q

		slices.SortFunc(b, f)
		q = b[:0]
		for _, v := range b {
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).detCmp(v.sub(q[len(q)-1])) >= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
		hulls[v][1] = q

		p := nodes[v]
		for _, e := range g[v] {
			w := e.to
			if w == fa {
				continue
			}
			down := p.fi
			if w == p.fiW {
				down = p.se
			}
			reroot(w, v, vec66{e.wt, max(up.x+up.y, down)})
		}
	}
	reroot(0, -1, vec66{})

	Fscan(in, &Q)
	p := vec66{0, 1}
	for range Q {
		var v int
		Fscan(in, &v, &p.x)
		h := hulls[v-1][0]
		j := sort.Search(len(h)-1, func(j int) bool { return p.dot(h[j]) > p.dot(h[j+1]) })
		mx := p.dot(h[j])
		mx2 := 0
		if j > 0 {
			mx2 = p.dot(h[j-1])
		}
		if j < len(h)-1 {
			mx2 = max(mx2, p.dot(h[j+1]))
		}
		h = hulls[v-1][1]
		if len(h) > 0 {
			j := sort.Search(len(h)-1, func(j int) bool { return p.dot(h[j]) > p.dot(h[j+1]) })
			mx2 = max(mx2, p.dot(h[j]))
		}
		Fprintln(out, max(mx+mx2, ans))
	}
}

//func main() { cf1866K(bufio.NewReader(os.Stdin), os.Stdout) }
