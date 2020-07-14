package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct{ l, r, min, max, cnt int }

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

func (t seg) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].min = min(lo.min, ro.min)
	t[o].max = max(lo.max, ro.max)
	t[o].cnt = lo.cnt + ro.cnt
}

func (t seg) _build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t._build(o<<1, l, m)
	t._build(o<<1|1, m+1, r)
}

func (t seg) _clear(o int) {
	if t[o].cnt == 0 {
		return
	}
	t[o].min = 0
	t[o].max = 0
	t[o].cnt = 0
	if t[o].l == t[o].r {
		return
	}
	t._clear(o << 1)
	t._clear(o<<1 | 1)
}

func (t seg) _update(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].min = v
		t[o].max = v
		t[o].cnt = 1
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t._update(o<<1, i, v)
	} else {
		t._update(o<<1|1, i, v)
	}
	t._pushUp(o)
}

func (t seg) _query(o, l, r, upp int) int {
	to := t[o]
	if l <= to.l && to.r <= r {
		if upp < to.min {
			return 0
		}
		if upp >= to.max {
			t._clear(o)
			return to.cnt
		}
	}
	res := 0
	m := (to.l + to.r) >> 1
	if l <= m {
		res += t._query(o<<1, l, r, upp)
	}
	if m < r {
		res += t._query(o<<1|1, l, r, upp)
	}
	t._pushUp(o)
	return res
}

func (t seg) init(n int)              { t._build(1, 1, n) }
func (t seg) update(i, v int)         { t._update(1, i, v) }
func (t seg) query(l, r, upp int) int { return t._query(1, l, r, upp) }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	buf := make([]byte, 4096)
	_i := len(buf)
	rc := func() byte {
		if _i == len(buf) {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n := r()
	t := make(seg, 4*n)
	t.init(n)
	for q := r(); q > 0; q-- {
		if r() == 1 {
			t.update(r()+1, r())
		} else {
			Fprintln(out, t.query(r()+1, r(), r()))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
