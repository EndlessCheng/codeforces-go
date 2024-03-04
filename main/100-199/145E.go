package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type data45 struct{ c0, c1, c01, c10 int }

type seg45 []struct {
	l, r int
	d    data45
	flip bool
}

func mergeInfo(a, b data45) (c data45) {
	c.c0 = a.c0 + b.c0
	c.c1 = a.c1 + b.c1
	c.c01 = max(a.c01+b.c1, a.c0+b.c01)
	c.c10 = max(a.c10+b.c0, a.c1+b.c10)
	return
}

func (t seg45) do(O int) {
	o := &t[O]
	d := o.d
	o.d = data45{d.c1, d.c0, d.c10, d.c01}
	o.flip = !o.flip
}

func (t seg45) build(s string, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		c1 := int(s[l-1] & 1)
		t[o].d = data45{c1 ^ 1, c1, 1, 1}
		return
	}
	m := (l + r) >> 1
	t.build(s, o<<1, l, m)
	t.build(s, o<<1|1, m+1, r)
	t[o].d = mergeInfo(t[o<<1].d, t[o<<1|1].d)
}

func (t seg45) update(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o)
		return
	}
	if t[o].flip {
		t.do(o << 1)
		t.do(o<<1 | 1)
		t[o].flip = false
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t[o].d = mergeInfo(t[o<<1].d, t[o<<1|1].d)
}

func cf145E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, l, r int
	var s string
	Fscan(in, &n, &m, &s)
	t := make(seg45, 2<<bits.Len(uint(n-1)))
	t.build(s, 1, 1, n)
	for ; m > 0; m-- {
		Fscan(in, &s)
		if s[0] == 's' {
			Fscan(in, &l, &r)
			t.update(1, l, r)
		} else {
			Fprintln(out, t[1].d.c01)
		}
	}
}

//func main() { cf145E(os.Stdin, os.Stdout) }
