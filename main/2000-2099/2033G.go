package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type seg33 []struct{ l, r, mx int }

func (t seg33) maintain(o int) {
	t[o].mx = max(t[o<<1].mx, t[o<<1|1].mx)
}

func (t seg33) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg33) update(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].mx = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg33) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mx
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf2033G(in io.Reader, out io.Writer) {
	var T, n, q int
	for Fscan(in, &T); T > 0; T-- {
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
		Fscan(in, &q)
		type pair struct{ k, i int }
		qs := make([][]pair, n)
		for i := range q {
			var v, k int
			Fscan(in, &v, &k)
			qs[v-1] = append(qs[v-1], pair{k, i})
		}

		type tuple struct{ fi, se, w int }
		downDis := make([]tuple, n)
		var build func(int, int)
		build = func(v, fa int) {
			fi, se, fw := 0, 0, -2
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				build(w, v)
				d := downDis[w].fi + 1
				if d > fi {
					se = fi
					fi, fw = d, w
				} else if d > se {
					se = d
				}
			}
			downDis[v] = tuple{fi, se, fw}
		}
		build(0, -1)

		maxD := downDis[0].fi
		t := make(seg33, 2<<bits.Len(uint(maxD)))
		t.build(1, 0, maxD)

		ans := make([]any, q)
		var dfs func(int, int, int)
		dfs = func(v, fa, d int) {
			for _, p := range qs[v] {
				if d == 0 || p.k == 0 {
					ans[p.i] = downDis[v].fi
				} else {
					ans[p.i] = max(t.query(1, max(d-p.k, 0), d-1)+d, downDis[v].fi)
				}
			}
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				mx := downDis[v].fi
				if w == downDis[v].w {
					mx = downDis[v].se
				}
				t.update(1, d, mx-d)
				dfs(w, v, d+1)
			}
		}
		dfs(0, -1, 0)
		Fprintln(out, ans...)
	}
}

//func main() { cf2033G(bufio.NewReader(os.Stdin), os.Stdout) }
