package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var a = []int{0}

type seg []struct {
	l, r, todo, ds, maxS int // 左闭右开
	leaf                 bool
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].ds = lo.ds + ro.ds
	t[o].maxS = max(lo.maxS, lo.ds+ro.maxS)
}

func (t seg) build(o, l, r int) {
	if l+1 == r {
		t[o].l = a[l-1]
		t[o].r = a[l] // l==r-1
		t[o].todo = 1e9 + 1
		t[o].leaf = true // 根节点必须标记
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m, r)
	t[o].l = t[o<<1].l // 从子节点获取
	t[o].r = t[o<<1|1].r
}

func (t seg) do(o, d int) {
	to := &t[o]
	to.todo = d
	to.ds = (to.r - to.l) * d
	to.maxS = max(to.ds, 0)
}

func (t seg) spread(o int) {
	if d := t[o].todo; d <= 1e9 {
		t.do(o<<1, d)
		t.do(o<<1|1, d)
		t[o].todo = 1e9 + 1
	}
}

func (t seg) update(o, l, r, d int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, d)
		return
	}
	t.spread(o)
	m := t[o<<1].r
	if l < m {
		t.update(o<<1, l, r, d)
	}
	if m < r {
		t.update(o<<1|1, l, r, d)
	}
	t.maintain(o)
}

func (t seg) query(o, sumUpp int) int {
	if t[o].leaf {
		l, r, ds := t[o].l, t[o].r, t[o].ds
		if ds <= 0 {
			return r
		}
		return min(r, l+sumUpp/(ds/(r-l)))
	}
	t.spread(o)
	if sumUpp < t[o<<1].maxS {
		return t.query(o<<1, sumUpp)
	}
	return t.query(o<<1|1, sumUpp-t[o<<1].ds)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type query struct{ l, r, d int }
	unique := func(a []int) []int {
		j, n := 0, len(a)
		for i := 1; i < n; i++ {
			if a[j] != a[i] {
				j++
				a[j] = a[i]
			}
		}
		return a[:j+1]
	}

	var sz, l, r, d int
	var op string
	Fscan(in, &sz)
	qs := []query{}
	for {
		Fscan(in, &op)
		if op[0] == 'E' {
			break
		}
		if op[0] == 'I' {
			Fscan(in, &l, &r, &d)
			qs = append(qs, query{l - 1, r, d}) // 也可以写成 l r+1，但是后面计算的时候涉及到 r 的时候还要减一，那不如写成 l-1 r
			a = append(a, l-1, r)
		} else {
			Fscan(in, &d)
			qs = append(qs, query{d: d})
		}
	}
	a = append(a, sz)

	sort.Ints(a)
	a = unique(a)
	n := len(a)
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for _, q := range qs {
		if q.r > 0 {
			t.update(1, q.l, q.r, q.d)
		} else {
			Fprintln(out, t.query(1, q.d))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
