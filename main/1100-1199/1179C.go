package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type seg79 []struct{ l, r, s, mx int }

func (t seg79) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg79) update(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].s += v
		t[o].mx += v
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].s = lo.s + ro.s
	t[o].mx = max(lo.mx+ro.s, ro.mx)
}

func (t seg79) binarySearch(o, s int) int {
	if t[o].l == t[o].r {
		if t[o].s > s {
			return t[o].l
		}
		return -1
	}
	if t[o<<1|1].mx > s {
		return t.binarySearch(o<<1|1, s)
	}
	return t.binarySearch(o<<1, s-t[o<<1|1].s)
}

func CF1179C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 1e6

	t := make(seg79, 4*mx)
	t.build(1, 1, mx)
	var n, m, q, op, p, v int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		t.update(1, a[i], 1) // 类似括号匹配，如此更新后，就可以从后往前找第一个大于 0 的位置了
	}
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
		t.update(1, b[i], -1)
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op, &p, &v)
		p--
		if op == 1 {
			t.update(1, a[p], -1)
			t.update(1, v, 1)
			a[p] = v
		} else {
			t.update(1, b[p], 1)
			t.update(1, v, -1)
			b[p] = v
		}
		Fprintln(out, t.binarySearch(1, 0))
	}
}

//func main() { CF1179C(os.Stdin, os.Stdout) }
