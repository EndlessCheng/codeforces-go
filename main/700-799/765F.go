package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type seg65 []struct {
	l, r, min int
	a         []int
}

func (t seg65) build(a []int, o, l, r int) {
	t[o].l, t[o].r, t[o].min = l, r, 2e9
	t[o].a = append([]int(nil), a[l-1:r]...)
	sort.Ints(t[o].a)
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
}

var curMin int

func (t seg65) update(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].min = min(t[o].min, abs65(v-t[o].a[0]))
		curMin = min(curMin, t[o].min)
		return
	}
	if t[o].r <= i {
		a := t[o].a
		p := sort.SearchInts(a, v)
		if (p == 0 || v-a[p-1] >= curMin) && (p == len(a) || a[p]-v >= curMin) {
			curMin = min(curMin, t[o].min)
			return // 提前退出
		}
	}
	m := (t[o].l + t[o].r) >> 1
	if i > m {
		t.update(o<<1|1, i, v) // 先右后左，这样可以从右往左更新
	}
	t.update(o<<1, i, v)
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg65) query(o, l int) int {
	if l <= t[o].l {
		return t[o].min
	}
	if (t[o].l+t[o].r)>>1 < l {
		return t.query(o<<1|1, l)
	}
	return min(t.query(o<<1, l), t[o<<1|1].min)
}

func CF765F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &q)
	qs := make([]struct{ l, r, i int }, q)
	for i := range qs {
		Fscan(in, &qs[i].l, &qs[i].r)
		qs[i].i = i
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i].r < qs[j].r })

	ans := make([]int, q)
	t := make(seg65, 4*n)
	t.build(a, 1, 1, n)
	for r, qi := 2, 0; r <= n; r++ {
		curMin = 2e9
		t.update(1, r-1, a[r-1])
		for ; qi < q && qs[qi].r == r; qi++ {
			ans[qs[qi].i] = t.query(1, qs[qi].l)
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func abs65(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//func main() { CF765F(os.Stdin, os.Stdout) }
