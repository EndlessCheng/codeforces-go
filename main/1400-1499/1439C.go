package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type seg39 []struct {
	l, r, todo, max, min int
	sum                  int64
}

func (t seg39) maintain(o int) {
	l, r := t[o<<1], t[o<<1|1]
	t[o].max = l.max
	t[o].min = r.min
	t[o].sum = l.sum + r.sum
}

func (t seg39) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].max = a[l-1]
		t[o].min = a[l-1]
		t[o].sum = int64(a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg39) do(O, v int) {
	o := &t[O]
	o.todo = v
	o.max = v
	o.min = v
	o.sum = int64(o.r-o.l+1) * int64(v)
}

func (t seg39) spread(o int) {
	if v := t[o].todo; v > 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg39) update(o, r, v int) {
	if v <= t[o].min {
		return
	}
	if t[o].r <= r && v >= t[o].max {
		t.do(o, v)
		return
	}
	t.spread(o)
	t.update(o<<1, r, v)
	if (t[o].l+t[o].r)>>1 < r {
		t.update(o<<1|1, r, v)
	}
	t.maintain(o)
}

func (t seg39) query(o, l int, v *int64) int {
	if *v < int64(t[o].min) {
		return 0
	}
	if l <= t[o].l && *v >= t[o].sum {
		*v -= t[o].sum
		return t[o].r - t[o].l + 1
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if m < l {
		return t.query(o<<1|1, l, v)
	}
	vl := t.query(o<<1, l, v)
	vr := t.query(o<<1|1, l, v)
	return vl + vr
}

func CF1439C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, tp, p, v int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg39, 4*n)
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &tp, &p, &v); tp == 1 {
			t.update(1, p, v)
		} else {
			v := int64(v)
			Fprintln(out, t.query(1, p, &v))
		}
	}
}

//func main() { CF1439C(os.Stdin, os.Stdout) }
