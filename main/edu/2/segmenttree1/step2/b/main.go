package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct{ l, r, sum int }

func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.init(a)
	return t
}

func (t seg) _pushUp(o int) { t[o].sum = t[o<<1].sum + t[o<<1|1].sum }

func (t seg) _build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _flip(o, idx int) {
	if t[o].l == t[o].r {
		t[o].sum ^= 1
		return
	}
	if idx <= (t[o].l+t[o].r)>>1 {
		t._flip(o<<1, idx)
	} else {
		t._flip(o<<1|1, idx)
	}
	t._pushUp(o)
}

func (t seg) _query(o, k int) int {
	if t[o].l == t[o].r {
		return t[o].l
	}
	if k < t[o<<1].sum {
		return t._query(o<<1, k)
	}
	return t._query(o<<1|1, k-t[o<<1].sum)
}

func (t seg) init(a []int)    { t._build(a, 1, 1, len(a)) }
func (t seg) flip(idx int)    { t._flip(1, idx) }
func (t seg) query(k int) int { return t._query(1, k) }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, v int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := newSegmentTree(a)
	for ; q > 0; q-- {
		if Fscan(in, &op, &v); op == 1 {
			t.flip(v + 1)
		} else {
			Fprintln(out, t.query(v)-1)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
