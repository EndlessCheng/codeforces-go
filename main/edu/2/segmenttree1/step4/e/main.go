package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct{ l, r, min int }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t seg) _pushUp(o int) { t[o].min = min(t[o<<1].min, t[o<<1|1].min) }

func (t seg) _build(o, l, r int) {
	t[o].l, t[o].r, t[o].min = l, r, 2e9
	if l == r {
		return
	}
	m := (l + r) >> 1
	t._build(o<<1, l, m)
	t._build(o<<1|1, m+1, r)
}

func (t seg) _update(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].min = v
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
	if l <= to.l && to.r <= r && upp < to.min {
		return 0
	}
	if to.l == to.r {
		t[o].min = 2e9
		return 1
	}
	defer t._pushUp(o)
	m := (to.l + to.r) >> 1
	if r <= m {
		return t._query(o<<1, l, r, upp)
	}
	if l > m {
		return t._query(o<<1|1, l, r, upp)
	}
	return t._query(o<<1, l, r, upp) + t._query(o<<1|1, l, r, upp)
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
