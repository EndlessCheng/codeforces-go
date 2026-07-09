package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type data struct{ state, mx int }
type tag struct{ flipCnt, add int }
type seg []struct {
	data
	tag
}

func mergeData(l, r data) data {
	state := 1
	if l.state == r.state {
		state = l.state
	}
	return data{state, max(l.mx, r.mx)}
}

func mergeTag(f, old tag) tag {
	if f.flipCnt == 0 {
		old.add += f.add
		return old
	}
	f.flipCnt += old.flipCnt
	return f
}

func (t seg) apply(o int, f tag) {
	cur := &t[o]
	if f.flipCnt > 0 {
		if f.flipCnt&1 > 0 {
			cur.state = 2 - cur.state
		}
		cur.mx = 0
	}
	if cur.state != 2 {
		cur.mx += f.add
	}

	cur.tag = mergeTag(f, cur.tag)
}

func (t seg) spread(o int) {
	f := t[o].tag
	if f == (tag{}) {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].tag = tag{}
}

func (t seg) update(o, l, r, ql, qr int, f tag) {
	if ql <= l && r <= qr {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		t.update(o<<1, l, m, ql, qr, f)
	}
	if m < qr {
		t.update(o<<1|1, m+1, r, ql, qr, f)
	}
	t[o].data = mergeData(t[o<<1].data, t[o<<1|1].data)
}

func (t seg) query(o, l, r, ql, qr int) int {
	if ql <= l && r <= qr {
		return t[o].mx
	}
	t.spread(o)
	m := (l + r) >> 1
	if qr <= m {
		return t.query(o<<1, l, m, ql, qr)
	}
	if ql > m {
		return t.query(o<<1|1, m+1, r, ql, qr)
	}
	return max(t.query(o<<1, l, m, ql, qr), t.query(o<<1|1, m+1, r, ql, qr))
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r, x int
	Fscan(in, &n, &q)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	for range q {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			Fscan(in, &x)
			t.update(1, 1, n, l, r, tag{0, x})
		} else if op == 2 {
			t.update(1, 1, n, l, r, tag{1, 0})
		} else {
			Fprintln(out, t.query(1, 1, n, l, r))
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
