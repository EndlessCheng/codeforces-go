package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type seg []struct {
	l, r, todo         int
	ans, pre, suf, sum int64
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].ans = max(max(lo.ans, ro.ans), lo.suf+ro.pre)
	t[o].pre = max(lo.pre, lo.sum+ro.pre)
	t[o].suf = max(ro.suf, ro.sum+lo.suf)
	t[o].sum = lo.sum + ro.sum
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r, t[o].todo = l, r, 1e9+1
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(o, v int) {
	to := &t[o]
	to.todo = v
	to.sum = int64(to.r-to.l+1) * int64(v)
	s := max(to.sum, 0)
	to.ans = s
	to.pre = s
	to.suf = s
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 1e9+1 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 1e9 + 1
	}
}

func (t seg) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r, v int
	Fscan(in, &n, &q)
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for ; q > 0; q-- {
		Fscan(in, &l, &r, &v)
		t.update(1, l+1, r, v)
		Fprintln(out, t[1].ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
