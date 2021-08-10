package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type seg57 []struct{ l, r, max, from, todoMax, todoFrom int }

func (t seg57) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg57) do(o, v, i int) {
	to := &t[o]
	if v > to.todoMax {
		to.todoMax, to.todoFrom = v, i
	}
	if v > to.max {
		to.max, to.from = v, i
	}
}

func (t seg57) spread(o int) {
	if v := t[o].todoMax; v != 0 {
		t.do(o<<1, v, t[o].todoFrom)
		t.do(o<<1|1, v, t[o].todoFrom)
		t[o].todoMax = 0
	}
}

func (t seg57) update(o, l, r, v, i int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v, i)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v, i)
	}
	if m < r {
		t.update(o<<1|1, l, r, v, i)
	}
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].max, t[o].from = max57(lo.max, lo.from, ro.max, ro.from)
}

func (t seg57) query(o, l, r int) (int, int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].max, t[o].from
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	a, af := t.query(o<<1, l, r)
	b, bf := t.query(o<<1|1, l, r)
	return max57(a, af, b, bf)
}

func CF1557D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ l, r int }

	var n, m, row, l, r int
	Fscan(in, &n, &m)
	rs := make([][]pair, n+1)
	x := make([]int, 0, m*2)
	for ; m > 0; m-- {
		Fscan(in, &row, &l, &r)
		rs[row] = append(rs[row], pair{l, r})
		x = append(x, l, r)
	}
	sort.Ints(x)
	kth, k := map[int]int{}, 1
	for i, v := range x {
		if i == 0 || v != x[i-1] {
			kth[v] = k
			k++
		}
	}

	t := make(seg57, k*4)
	t.build(1, 1, k)
	fs := make([]int, n+1)
	for i, rs := range rs {
		if rs == nil {
			continue
		}
		mx, from := 0, 0
		for _, p := range rs {
			v, f := t.query(1, kth[p.l], kth[p.r])
			if v > mx {
				mx, from = v, f
			}
		}
		fs[i] = from
		mx++
		for _, p := range rs {
			t.update(1, kth[p.l], kth[p.r], mx, i)
		}
	}
	Fprintln(out, n-t[1].max)
	save := make([]bool, n+1)
	for i := t[1].from; i > 0; i = fs[i] {
		save[i] = true
	}
	for i := 1; i <= n; i++ {
		if !save[i] {
			Fprint(out, i, " ")
		}
	}
}

func max57(a, af, b, bf int) (int, int) {
	if a > b {
		return a, af
	}
	return b, bf
}

//func main() { CF1557D(os.Stdin, os.Stdout) }
