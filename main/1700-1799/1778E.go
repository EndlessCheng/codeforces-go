package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type xorBasis78 struct{ b, pos [30]int }

func (b *xorBasis78) insertRightMost(v, p int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 {
			b.b[i] = v
			b.pos[i] = p
			return
		}
		if p >= b.pos[i] {
			p, b.pos[i] = b.pos[i], p
			v, b.b[i] = b.b[i], v
		}
		v ^= b.b[i]
	}
}

func (b *xorBasis78) maxXor(l int) (xor int) {
	for i := len(b.b) - 1; i >= 0; i-- {
		if xor>>i&1 == 0 && b.pos[i] >= l && xor^b.b[i] > xor {
			xor ^= b.b[i]
		}
	}
	return
}

func CF1778E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w, q, rt int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		dfnVal := make([]int, n)
		nodes := make([]struct{ dfn, size int }, n)
		dfn := -1
		const mx = 18
		pa := make([][mx]int, n)
		dep := make([]int, n)
		var build func(v, p, d int) int
		build = func(v, p, d int) int {
			pa[v][0] = p
			dep[v] = d
			dfn++
			nodes[v].dfn = dfn
			dfnVal[dfn] = a[v]
			sz := 1
			for _, w := range g[v] {
				if w != p {
					sz += build(w, v, d+1)
				}
			}
			nodes[v].size = sz
			return sz
		}
		build(0, -1, 0)
		for i := 0; i+1 < mx; i++ {
			for v := range pa {
				if p := pa[v][i]; p != -1 {
					pa[v][i+1] = pa[p][i]
				} else {
					pa[v][i+1] = -1
				}
			}
		}
		down := func(v, to int) int {
			if dep[v] >= dep[to] {
				return -1
			}
			d := dep[v] + 1
			for i := 0; i < mx; i++ {
				if (dep[to]-d)>>i&1 > 0 {
					to = pa[to][i]
				}
			}
			if pa[to][0] == v {
				return to
			}
			return -1
		}

		Fscan(in, &q)
		type query struct{ l, i int }
		qs := make([][]query, n*2)
		for i := 0; i < q; i++ {
			Fscan(in, &rt, &v)
			if rt == v {
				qs[n-1] = append(qs[n-1], query{0, i})
				continue
			}
			rt--
			v--
			d := down(v, rt)
			var l, r int
			if d < 0 {
				o := nodes[v]
				l, r = o.dfn, o.dfn+o.size-1
			} else {
				o := nodes[d]
				l, r = o.dfn+o.size, o.dfn+n-1
			}
			qs[r] = append(qs[r], query{l, i})
		}

		ans := make([]int, q)
		b := &xorBasis78{}
		for r, qs := range qs {
			b.insertRightMost(dfnVal[r%n], r)
			for _, q := range qs {
				ans[q.i] = b.maxXor(q.l)
			}
		}
		for _, v := range ans {
			Fprintln(out, v)
		}
	}
}

//func main() { CF1778E(os.Stdin, os.Stdout) }
