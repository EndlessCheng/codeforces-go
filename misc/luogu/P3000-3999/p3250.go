package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
type fenwick250 []int
var rec250 []int

func (t fenwick250) reset() {
	for _, i := range rec250 {
		for ; i < len(t); i += i & -i {
			t[i] = 0
		}
	}
	rec250 = rec250[:0]
}

func (t fenwick250) update(i, val int) {
	rec250 = append(rec250, i)
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick250) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func (t fenwick250) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func p3250(in io.Reader, _w io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for range n - 1 {
		var x, y int
		Fscan(in, &x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	tin := make([]int, n+1)
	tout := make([]int, n+1)
	dfn := 0
	const mx = 17
	pa := make([][mx]int, n+1)
	dep := make([]int, n+1)
	var dfs func(int, int)
	dfs = func(x, p int) {
		dfn++
		tin[x] = dfn
		pa[x][0] = p
		for i := range mx - 1 {
			pa[x][i+1] = pa[pa[x][i]][i]
		}
		for _, y := range g[x] {
			if y != p {
				dep[y] = dep[x] + 1
				dfs(y, x)
			}
		}
		tout[x] = dfn
	}
	dfs(1, 0)

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
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}

	type query struct{ a, b, v int }
	qs := make([]query, m)
	idx := make([]int, m)
	sorted := make([]int, 0, m)
	for i := range qs {
		var op, a, b, v int
		Fscan(in, &op, &a)
		if op == 0 {
			Fscan(in, &b, &v)
			qs[i] = query{a, b, v}
			sorted = append(sorted, v)
		} else if op == 1 {
			qs[i] = qs[a-1]
			qs[i].v *= -1
		} else {
			qs[i] = query{-a, -1, 0}
		}
		idx[i] = i
	}

	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	t := make(fenwick250, n+1)

	var solve func([]int, int, int)
	solve = func(idx []int, low, high int) {
		for _, i := range idx {
			if qs[i].a < 0 {
				goto next
			}
		}
		return

	next:
		if low+1 == high { // 开区间为空
			if low >= 0 {
				for _, i := range idx {
					if qs[i].a < 0 {
						qs[i].b = sorted[low]
					}
				}
			}
			return
		}

		mid := (low + high) >> 1
		midVal := sorted[mid]
		var b, c []int

		pathCnt := 0
		for _, i := range idx {
			q := qs[i]
			if q.a > 0 { // 修改
				x, y, v := q.a, q.b, q.v
				if abs(v) < midVal {
					b = append(b, i)
					continue
				}
				if v > 0 {
					v = 1
				} else {
					v = -1
				}
				pathCnt += v
				t.update(tin[x], v)
				t.update(tin[y], v)
				lca := getLCA(x, y)
				t.update(tin[lca], -v)
				if f := pa[lca][0]; f > 0 {
					t.update(tin[f], -v)
				}
				c = append(c, i)
			} else { // 查询
				x := -q.a
				if t.query(tin[x], tout[x]) == pathCnt {
					b = append(b, i)
				} else {
					c = append(c, i)
				}
			}
		}

		// 撤销修改
		t.reset()

		solve(b, low, mid)
		solve(c, mid, high)
	}

	// 开区间二分
	solve(idx, -1, len(sorted))
	for _, q := range qs {
		if q.a < 0 {
			Fprintln(out, q.b)
		}
	}
}

//func main() { p3250(bufio.NewReader(os.Stdin), os.Stdout) }
